package mail

import (
	"testing"
)

func TestMail(t *testing.T) {
	s := Smtp {
		Host: "smtp.gmail.com",
		Port: "587",
		User: "greenkiweez1@gmail.com",
		Pass: "22222",
	}
	to := []string{"kristjan.voje@gmail.com"}
	msg := []byte("Subject: testing something\nSomeBodyOnceToldMe")
	s.SendEmail(to, msg)
}
