package util

import (
	"bytes"
	"html/template"
	"net/smtp"
	"log"
	"strconv"
	"github.com/ilm-statistics/ilm-statistics/model"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/scorredoira/email"
	"net/mail"
)

const (
	EMAILCONFIGFILE = "./emailConfig.yml"
	TEMPLATEPATH = "./processor/util/emailTemplate.html"
)

type EmailConfig struct {
	Subject string
	SmtpServer string
	SmtpPort string
	From string
	Password string
	To []string
}

type StatsToSend struct {
	Ip string
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
	MostPopularProjects []model.Project
	MaxProjectPopularity int
	ImagesInProjects map[string][]model.Project
	ProjectsList map[string]model.Project
	ProjectsSuccess map[string]string
	ProjectsFailure map[string]string
	MostUsedImages model.PairList
	LeastUsedImages []string
	LeastUsedImageOccurrence int
	MostExecutedTests []model.Test
	MostExecutedTestsNr int
	LeastExecutedTests []model.Test
	LeastExecutedTestsNr int
	Vulnerabilities model.NoOfVulnerabilitiesWithLinksList
	ImagesInRegistries map[string][]string
}

type StatsToSendList []StatsToSend

var emailConfig EmailConfig

func SendEmailTemplate(stat model.Statistic, statForIp map[string]model.Statistic) {
	log.Println("Start initializing the e-mail")

	log.Println("Loading e-mail configurations from file")
	ParseEmailConfigFile()

	statForIp[""] = stat

	templateData := ConvertStatToTemplate(statForIp)

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

	addr := emailConfig.SmtpServer + ":" + emailConfig.SmtpPort

	m := email.NewHTMLMessage(r.subject, r.body)
	m.From = mail.Address{Address: emailConfig.From}
	m.To = emailConfig.To

	auth := smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.SmtpServer)
	if err := email.Send(addr, auth, m); err != nil {
		log.Fatal(err)
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

func ConvertStatToTemplate(stats map[string]model.Statistic) StatsToSendList {
	data := []StatsToSend{}

	for ip, stat := range stats {
		templateData := StatsToSend{
			Ip: ip,
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
			ProjectsList: stat.Projects.IdToProject,
			MostUsedImages: stat.MostUsedImages,
			MostExecutedTests: stat.MostExecutedTests,
			MostExecutedTestsNr: stat.MostExecutedTestsNr,
			LeastExecutedTests: stat.LeastExecutedTests,
			LeastExecutedTestsNr: stat.LeastExecutedTestsNr,
			Vulnerabilities: stat.Vulnerabilities,
			ImagesInRegistries: stat.RegistriesAndImages,
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

		if ip != "" {
			data = append(data, templateData) // append
		} else {
			data = append([]StatsToSend{templateData}, data...) //prepend
		}
	}

	return data
}