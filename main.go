package main

import (
	"email-sender/m/src"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func emailList () ([]string, error) {
	list :=  os.Getenv("TARGETLIST")
	if list == "" {
		return nil, errors.New("TARGETLIST is mandatory")
	}
	emailList := strings.Split(list, ",")
	return emailList, nil
}

func main() {
	env := src.EnvStruct{}
	data, e := env.Verify()
	if e != nil {
		panic(e.Error())
	}

	smtpServer := src.SmtpServerStruct{
		Host: data.SmtpHost,
		Port: data.SmtpPort,
	}

	message := []byte("hi")

	auth := smtp.PlainAuth("", data.SmtpUser, data.SmtpPassWord, smtpServer.Host)

	err := smtp.SendMail(
		smtpServer.ServerName(), auth, data.From, data.To, message,
		)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

