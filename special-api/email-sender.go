package main

import (
	"net/smtp"
	//	"crypto/tls"
	"fmt"
	//"github.com/sendgrid/sendgrid-go"
	//"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func (e *Email) BuildMessage() []byte {
	m := "Subject :" + e.Subject + "\n \n" + e.Body
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

func main() {
	email := Email{
		From:    "nader.special.api@gmail.com",
		To:      []string{"nader_atef80@outlook.com"},
		Subject: "hello",
		Body:    "hello its me \n hello \n hello",
	}
	server := SmtpServer{
		Host: "smtp.gmail.com",
		Port: "587",
	}
	//host := server.ReturnHostName
	auth := smtp.PlainAuth("", email.From, "password", server.Host)
	err := smtp.SendMail(server.ReturnHostName(), auth, email.From, email.To, email.BuildMessage())
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent, ")
}
