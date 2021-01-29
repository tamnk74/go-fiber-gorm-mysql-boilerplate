package Email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"
	"time"

	"github.com/tamnk74/todolist-mysql-go/config"
)

func Send(to []string, subject string) {
	// Authentication.
	auth := smtp.PlainAuth(config.MAIL.FROM, config.MAIL.USERNAME, config.MAIL.PASSWORD, config.MAIL.HOST)

	t, _ := template.ParseFiles("templates/email/welcome.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s \n%s\n\n", config.MAIL.FROM, strings.Join(to, ", "), subject, mimeHeaders)))

	t.Execute(&body, struct {
		Name string
	}{
		Name: "Tom NK",
	})

	// Sending email.
	err := smtp.SendMail(config.MAIL.HOST+":"+config.MAIL.PORT, auth, config.MAIL.FROM, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("An email has been sent from ", config.MAIL.FROM, " to ", strings.Join(to, ", "), " at ", time.Now())
}
