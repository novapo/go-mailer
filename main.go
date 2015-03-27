package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"text/template"
)

type mailParams struct {
	From    string
	To      string
	Name    string
	Email   string
	Message string
}

type formData struct {
	Name    string `json:"c_name"`
	Email   string `json:"c_email"`
	Message string `json:"c_message"`
}

type response struct {
	Status  int    `json:"sendstatus"`
	Message string `json:"message"`
}

func main() {
	usage()

	err := http.ListenAndServe(":"+strconv.Itoa(port), http.HandlerFunc(handler))
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	formData := &formData{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	parameters := mailParams{username, recipients.String(), formData.Name, formData.Email, formData.Message}
	response := &response{}

	auth := smtp.PlainAuth("", username, password, smtpHost)

	buffer := new(bytes.Buffer)
	template := template.Must(template.New("emailTemplate").Parse(emailScript()))
	template.Execute(buffer, &parameters)

	err = smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, username, recipients, buffer.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		response.Status = 0
		response.Message = "Email konnte nicht versendet werden."
	} else {
		w.WriteHeader(http.StatusOK)
		response.Status = 1
		response.Message = "Vielen Dank für Ihre Anfrage."
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func emailScript() (script string) {
	return `From: {{.From}}
To: {{.To}}
Subject: Anfrage von {{.Name}}({{.Email}})
MIME-version: 1.0
Content-Type: text/plain; charset="UTF-8"

{{.Message}}`
}
