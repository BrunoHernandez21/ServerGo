package moduls

import (
	"cursos_app/src/config"
	"net/smtp"
)

func Send_Mail_Password(email string, password string) bool {

	from := smtp.PlainAuth("", config.Mail.Email, config.Mail.Password, config.Mail.Host)
	to := []string{email}

	titulo := "From: Aldo Nery "
	Subtitulo := "Nueva contraseña "
	msg := []byte(titulo + "<" + config.Mail.Email + ">\n" + "Subject: " + Subtitulo + "\n" + "\nSu nueva contraseña se muestra a continuacion\nPassword: " + password)
	error := smtp.SendMail(config.Mail.ServerName, from, config.Mail.Email, to, msg)
	return error != nil
}
