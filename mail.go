package message

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Mail struct {
	Host     string
	Protocol string
	Port     int
	Username string
	Password string
	FromName string
	To       []string
}

func (m *Mail) Send(body Body) error {
	//log.Printf("[mail] sending message...")
	e := email.NewEmail()
	e.From = m.FromName
	e.To = m.To
	e.Subject = body.Title
	e.Text = []byte(body.Content)
	addr := fmt.Sprintf("%v:%v", m.Host, m.Port)
	err := e.Send(addr, smtp.PlainAuth("", m.Username, m.Password, m.Host))
	if err != nil {
		//log.Fatalf("[mail] send failed: %v\n", err)
		return err
	}
	//log.Printf("[mail] send successful")
	return nil
}
