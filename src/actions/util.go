package actions

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/brenddonanjos/go_api/src/models"
)

func SendMail(msg string) error {
	// Sender data.
	from := "<Email>"
	password := "<Password>"

	// Receiver email address.
	to := []string{
		"brenddon.dev@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return err
}

func ErrControl(err error) {
	SendMail(err.Error()) //send warnning mail
	l := Log{
		Message:   err.Error(),
		Type:      "error",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	l.NewLog() //save on table logs
}

func ToSlice(c <-chan models.Article) []models.Article {
	s := models.Articles{}
	for i := range c {
		s = append(s, i)
	}
	return s
}
