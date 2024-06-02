package mail

import (
	"emailn/internal/config"
	"emailn/internal/interfaces"
	"gopkg.in/gomail.v2"
)

func SendMail(params interfaces.MailParams) error {
	env := config.LoadEnv()

	sender := gomail.NewDialer(env.SMTPHost, 587, env.MailUser, env.MailPassword)
	message := gomail.NewMessage()
	message.SetHeader("From", env.MailUser)
	message.SetHeader("To", params.Email)
	message.SetHeader("Subject", params.Subject)
	message.SetBody("text/html", "<p>"+params.Content+"</p>")

	return sender.DialAndSend(message)
}
