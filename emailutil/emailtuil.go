package emailutil

import (
	"fmt"
	"strings"

	"github.com/go-gomail/gomail"
)

type MailUser struct {
	Email string
	Name  string
}

var (
	#mailServerHost = "smtp.exmail.qq.com"
	mailServerHost = ""
	mailServerPort = 465
)

var maintainers = []MailUser{
	{
		Email: "airdb@qq.com",
		Name:  "airdb",
	},
}

func SendEmail(toEmails, subject, content string) {
	m := gomail.NewMessage()

	sender := MailUser{
		Email: "dean@baobeihuijia.com",
		Name:  "Hello",
	}

	ccList := []string{}
	for _, maintainer := range maintainers {
		ccList = append(ccList, m.FormatAddress(maintainer.Email, maintainer.Name))
	}

	m.SetAddressHeader("From", sender.Email, sender.Name)

	toList := []string{}
	for _, toEmail := range strings.Split(toEmails, ",") {
		toList = append(toList, m.FormatAddress(toEmail, ""))
	}
	// m.SetHeader("To", m.FormatAddress(toMail.Email, toMail.Name))
	m.SetHeader("To", toList...)

	m.SetHeader("Cc", ccList...)

	// subject := fmt.Sprintf("%s[Space CertMS] %s", os.Getenv("ENV"), domain)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", content)

	d := gomail.NewDialer(mailServerHost, mailServerPort, sender.Email, "")
	fmt.Println("dd", d)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("err", err.Error())
		return
	}
}
