package moduls

import (
	"net/smtp"
)

func Send_Mail_Password(email string) bool {
	mainEmail := "ichimar21@gmail.com"
	password := "xpxqhixlrsldedzq"
	host := "smtp.gmail.com"
	serverName := "smtp.gmail.com:587"
	from := smtp.PlainAuth("", mainEmail, password, host)
	to := []string{email}

	titulo := "From: Aldo Nery "
	Subtitulo := "Nueva contraseña "
	msg := []byte(titulo + "<" + mainEmail + ">\n" + "Subject: " + Subtitulo + "\n" + "\nSu nueva contraseña se muestra a continuacion\nPassword: " + "123das")
	error := smtp.SendMail(serverName, from, "ichimar21@gmail.com", to, msg)
	return error != nil
}
