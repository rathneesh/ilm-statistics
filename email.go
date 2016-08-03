package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"log"
	"strconv"
)

var auth smtp.Auth

func SendEmailTemplate(stat Statistic) {
	log.Println("Start initializing e-mail template")
	templateData := struct {
		Users int
		Accounts int
		AverageAccountsPerUser string
		NrProj0 int
		NrProj1 int
		NrProj2 int
		NrProj3 int
		NrProj4 int
		NrProj5 int
		NrProj6 int
		NrProj7 int
		NrProj8 int
		NrProj9 int
		NrProj10 int
		NrProj11 int
		NrProj12 int
		NrProj13 int
		NrProj14 int
		NrProj15 int
		NrProj16 int
		NrProj17 int
		NrProj18 int
		NrProj19 int
		NrProj20 int
		NrProj21 int
		NrProj22 int
		NrProj23 int
		BusiestHour string
		Tests int
		Projects int
		AverageTestsPerProject string
		Images int
		AverageImagesPerProject string
		SuccessRate string
		FailureRate string
	}{
		Users: stat.Users,
		Accounts: stat.Accounts,
		AverageAccountsPerUser: strconv.FormatFloat(stat.AvgAccountPerUser, 'f', 2, 64),
		NrProj0: 0,
		NrProj1: 0,
		NrProj2: 0,
		NrProj3: 0,
		NrProj4: 0,
		NrProj5: 0,
		NrProj6: 0,
		NrProj7: 0,
		NrProj8: 0,
		NrProj9: 0,
		NrProj10: 0,
		NrProj11: 0,
		NrProj12: 0,
		NrProj13: 0,
		NrProj14: 0,
		NrProj15: 0,
		NrProj16: 0,
		NrProj17: 0,
		NrProj18: 0,
		NrProj19: 0,
		NrProj20: 0,
		NrProj21: 0,
		NrProj22: 0,
		NrProj23: 0,
		BusiestHour: "",
		Tests: stat.Tests.Total,
		Projects: stat.Projects.Total,
		AverageTestsPerProject: strconv.FormatFloat(stat.Projects.AvgTestsInProjects, 'f', 2, 64),
		Images: stat.Projects.ImagesInProjects,
		AverageImagesPerProject: strconv.FormatFloat(stat.Projects.AvgImagesInProjects, 'f', 2, 64),
		SuccessRate: strconv.FormatFloat(stat.Projects.SuccessRate, 'f', 2, 64),
		FailureRate: strconv.FormatFloat(stat.Projects.FailureRate, 'f', 2, 64),
	}
	log.Println("Creating new e-mail request")
	r := NewRequest([]string{"soninob@hpe.com", "lilla.vass@hpe.com"}, "[TEST] ILM Statistics", "Hello, World!")
	err := r.ParseTemplate("emailTemplate.html", templateData)
	if err != nil {
		log.Println(err)
		log.Println("Template could not be parsed, sending plain e-mail")
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
	r.SendEmail()
	log.Println("Sending the e-mail")
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
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp3.hpe.com:25"

	err := smtp.SendMail(addr, nil, "ilm.stats@hpe.com", r.to, msg)
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