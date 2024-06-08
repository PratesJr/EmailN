package mail

import (
	"emailn/internal/config"
	"emailn/internal/domain/campaign"
	"gopkg.in/gomail.v2"
)

func SendMail(params campaign.Campaign) error {
	env := config.LoadEnv()

	var mailList []string

	for _, contact := range params.Contacts {
		mailList = append(mailList, contact.Email)
	}
	sender := gomail.NewDialer(env.SMTPHost, 587, env.MailUser, env.MailPassword)
	message := gomail.NewMessage()
	message.SetHeader("From", env.MailUser)
	message.SetHeader("To", mailList...)
	message.SetHeader("Subject", params.Name)
	message.SetBody("text/html", "<p>"+params.Content+"</p>")

	return sender.DialAndSend(message)
}
