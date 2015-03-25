package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/novapo/go-mailer/config"
)

type MailParams struct {
	From    string
	To      string
	Subject string
	Name    string
	Email   string
	Message string
}

var (
	conf *config.Config
)

func main() {

	c, err := config.FromFile("config.json")

	if err != nil {
		log.Fatal(err)
	}

	conf = c

	err = http.ListenAndServe(conf.Addr, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	r.ParseForm()

	name := ""
	email := ""
	message := ""

	if tmp := r.Form[conf.FormData.Name]; len(tmp) != 0 {
		name = tmp[0]
	}

	if tmp := r.Form[conf.FormData.Email]; len(tmp) != 0 {
		email = tmp[0]
	}

	if tmp := r.Form[conf.FormData.Message]; len(tmp) != 0 {
		message = tmp[0]
	}

	if name == "" || email == "" || message == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	parameters := MailParams{conf.From, strings.Join([]string(conf.To), ","), conf.Subject, name, email, message}

	// Set up authentication information.
	auth := smtp.PlainAuth("", conf.Smtp.Username, conf.Smtp.Password, conf.Smtp.Host)

	buffer := new(bytes.Buffer)
	template := template.Must(template.New("emailTemplate").Parse(emailScript()))
	template.Execute(buffer, &parameters)

	fmt.Println(buffer)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	fmt.Println("Send mail...")
	err := smtp.SendMail(conf.Smtp.Host+":"+conf.Smtp.Port, auth, conf.From, conf.To, buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
	w.WriteHeader(http.StatusNoContent)
}

func emailScript() (script string) {
	return `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"

<h1>Anfrage von {{.Name}}({{.Email}})</h1>

{{.Message}}`
}
