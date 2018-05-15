package main

import (
	"net/smtp"
	//	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func (e *Email) BuildMessage() []byte {
	m := "Subject :" + e.Subject + "\n" + "from: " + e.address + "\n" + e.name + "\n" + e.Body
	fmt.Println(e)
	return []byte(m)
}

type SmtpServer struct {
	Host string
	Port string
}

func (s *SmtpServer) ReturnHostName() string {
	host := s.Host + ":" + s.Port
	return host
}

func sendEmail(w http.RepsoneWriter, name, address, message, pass string) {
	email := Email{
		From:    "nader.special.api@gmail.com",
		To:      []string{"nader_atef80@outlook.com"},
		Subject: "Email from portfolio site",
		Body:    message,
		address: address,
		name:    name,
	}
	server := SmtpServer{
		Host: "smtp.gmail.com",
		Port: "587",
	}
	//host := server.ReturnHostName
	auth := smtp.PlainAuth("", email.From, pass, server.Host)
	err := smtp.SendMail(server.ReturnHostName(), auth, email.From, email.To, email.BuildMessage())
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent, ")
}

func main() {
	pass := os.Getenv("gmailapipass")
	http.HandleFunc("/send", func(w http.ResponseWriter, r http.Request) {

	})
}
