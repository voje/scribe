package mail

import (
	"net/smtp"
)

type Smtp struct {
	Host string
	Port string
	User string
	Pass string
}

func (s *Smtp) SendEmail(to []string, msg []byte) error {
	auth := smtp.PlainAuth("", s.User, s.Pass, s.Host)
	addr := s.Host + ":" + s.Port
	err := smtp.SendMail(addr, auth, s.User, to, msg)
	return err
}
