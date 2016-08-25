package util

import (
	"bytes"
	"html/template"
	"net/smtp"
	"log"
	"strconv"
	"github.com/ilm-statistics/ilm-statistics/model"
	"strings"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	EMAILCONFIGFILE = "./emailConfig.yml"
	TEMPLATEPATH = "./processor/util/emailTemplate.html"
)

type EmailConfig struct {
	Subject string
	SmptServer string
	From string
	To []string
}

var emailConfig EmailConfig

func SendEmailTemplate(stat model.Statistic) {
	log.Println("Start initializing the e-mail")

	log.Println("Loading e-mail configurations from file")
	ParseEmailConfigFile()
	templateData := struct {
		Users int
		Accounts int
		AverageAccountsPerUser string
		IsActivity bool
		Hours map[int]int
		BusiestHours []int
		Tests int
		Projects int
		AverageTestsPerProject string
		Images int
		AverageImagesPerProject string
		SuccessRate string
		FailureRate string
		MostPopularProjects map[string]model.Project
		MaxProjectPopularity int
		ImagesInProjects map[string][]model.Project
		ProjectsList []model.Project
		ProjectsSuccess map[string]string
		ProjectsFailure map[string]string
		MostUsedImages []string
		MostUsedImageOccurrence int
		LeastUsedImages []string
		LeastUsedImageOccurrence int
		NumberOfImages int
		MostExecutedTests []model.Test
		MostExecutedTestsNr int
		LeastExecutedTests []model.Test
		LeastExecutedTestsNr int
	}{
		Users: stat.Users,
		Accounts: stat.Accounts,
		AverageAccountsPerUser: strconv.FormatFloat(stat.AvgAccountPerUser, 'f', 2, 64),
		IsActivity: false,
		Hours: stat.HourlyActivities,
		BusiestHours: stat.BusiestHours,
		Tests: stat.Tests.Total,
		Projects: stat.Projects.Total,
		AverageTestsPerProject: strconv.FormatFloat(stat.Projects.AvgTestsInProjects, 'f', 2, 64),
		Images: stat.Projects.ImagesInProjects,
		AverageImagesPerProject: strconv.FormatFloat(stat.Projects.AvgImagesInProjects, 'f', 2, 64),
		SuccessRate: strconv.FormatFloat(stat.Projects.SuccessRate, 'f', 2, 64),
		FailureRate: strconv.FormatFloat(stat.Projects.FailureRate, 'f', 2, 64),
		MostPopularProjects: stat.MostPopularProjects,
		MaxProjectPopularity: stat.MaxProjectPopularity,
		ImagesInProjects: stat.ImagesInProjects,
		ProjectsList: stat.ScriptProjects,
		MostUsedImages: stat.MostUsedImages,
		MostUsedImageOccurrence: stat.MostUsedImageOccurrence,
		LeastUsedImages: stat.LeastUsedImages,
		LeastUsedImageOccurrence: stat.LeastUsedImageOccurrence,
		NumberOfImages: stat.NumberOfImages,
		MostExecutedTests: stat.MostExecutedTests,
		MostExecutedTestsNr: stat.MostExecutedTestsNr,
		LeastExecutedTests: stat.LeastExecutedTests,
		LeastExecutedTestsNr: stat.LeastExecutedTestsNr,

	}

	if len(templateData.Hours) != 0 {
		templateData.IsActivity = true
	}


	templateData.ProjectsSuccess = map[string]string{}
	templateData.ProjectsFailure = map[string]string{}
	for projectId, success := range stat.ProjectsSuccess {
		templateData.ProjectsSuccess[projectId] = strconv.FormatFloat(success, 'f', 2, 64)
	}
	for projectId, failure := range stat.ProjectsFailure {
		templateData.ProjectsFailure[projectId] = strconv.FormatFloat(failure, 'f', 2, 64)
	}

	log.Println("Creating new e-mail request")
	r := NewRequest(emailConfig.To, emailConfig.Subject, "")
	err := r.ParseTemplate(TEMPLATEPATH, templateData)
	if err != nil {
		log.Println(err)
		log.Println("Template could not be parsed")
	} else {
		log.Println("Sending the e-mail")
		_, e := r.SendEmail()
		if e != nil {
			log.Println(e)
		}
	}

}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := MIME
	subject := "Subject: " + r.subject + "\n"
	to := "To: " + strings.TrimSuffix(strings.Join(emailConfig.To, ","), ",")
	msg := []byte(subject + "From: " + emailConfig.From + "\n" + to + "\n;" + mime + "\n\r" + r.body)
	addr := emailConfig.SmptServer

	err := smtp.SendMail(addr, nil, emailConfig.From, r.to, msg)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	log.Println("Start parsing the e-mail template")
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Println(err)
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Println(err)
		return err
	}
	r.body = buf.String()
	log.Println("End parsing the e-mail template")
	return nil
}

func ParseEmailConfigFile() {
	log.Println("Parsing the config file")
	filename, _ := filepath.Abs(EMAILCONFIGFILE)

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &emailConfig)

	if err != nil {
		log.Println(err)
		return
	}
}