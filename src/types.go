package src

import (
	"errors"
	"os"
	"strings"
)

type SmtpServerStruct struct {
	Host string
	Port string
}

func (s *SmtpServerStruct) ServerName () string {
	host := s.Host + ":" + s.Port
	return host
}

type EnvStruct struct {
	SmtpHost string
	SmtpPort string
	SmtpUser string
	SmtpPassWord string
	From string
	To []string
}

func (e *EnvStruct) Verify () (*EnvStruct, error) {
	smtpHost := os.Getenv("SMTP_HOSt")
	smtpPort := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("FROM")
	to := os.Getenv("TO")

	if smtpHost == "" {
		return nil, errors.New("SMTP_HOST is required")
	}
	if smtpPort == "" {
		return nil, errors.New("SMTP_PORT is required")
	}
	if user == "" {
		return nil, errors.New("SMTP_USER is required")
	}
	if password == "" {
		return nil, errors.New("SMTP_PASSWORD is required")
	}
	if from == "" {
		return nil, errors.New("FROM is required")
	}
	if to == "" {
		return nil, errors.New("TO is required")
	}
	e.SmtpHost = smtpHost
	e.SmtpPort = smtpPort
	e.SmtpUser = user
	e.SmtpPassWord = password
	e.From = from
	e.To = strings.Split(to, ",")
	return e, nil
}