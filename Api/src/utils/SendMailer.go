package utils

import (
	"fmt"
	env "main/src/envirimoents"
	"net/smtp"
)

func SendMailer(subject string, body string, to []string) bool {
	// Configura los parámetros del servidor de correo saliente (SMTP)
	smtpServer := env.GetEnv("SMTPSERVER", "smtp.example.com")
	smtpPort := env.GetEnv("SMTPPORT", "587")
	smtpUsername := env.GetEnv("SMTPUSERNAME", "user")
	smtpPassword := env.GetEnv("SMTPPASSWORD", "password")

	// Configura el mensaje de correo electrónico
	from := env.GetEnv("SMTPFROMEMAIL", "email@example.com")
	// to := []string{"dscampaz3110@gmail.com"}
	// subject := "Asunto del correo"
	// body := "Cuerpo del correo electrónico"

	tos := ""

	for _, value := range to {
		tos = tos + value
	}

	// Construye el mensaje de correo
	message := "Subject: " + subject + "\r\n" +
		"From: " + from + "\r\n" +
		"To: " + tos + "\r\n" +
		"\r\n" + body

	/*
	   Mensage con HTML
	   	message := "Subject: " + subject + "\r\n" +
	   		"From: " + from + "\r\n" +
	   		"To: " + to[0] + "\r\n" +
	   		"MIME-Version: 1.0\r\n" +
	   		"Content-Type: text/html; charset=\"utf-8\"\r\n" +
	   		"\r\n" + htmlBody

	*/

	// Establece la conexión con el servidor SMTP
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+fmt.Sprint(smtpPort), auth, from, to, []byte(message))
	if err != nil {
		return false
	}

	return true
}

func ErrorEmail(subj string, body string) bool {
	return SendMailer("Error to "+subj, body, []string{env.GetEnv("SMTPFROMEMAIL", "email@example.com")})
}
