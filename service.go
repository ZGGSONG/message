package message

import (
	"errors"
	"github.com/zggsong/message/internal/message"
)

type Service struct {
	Body message.Body
}

func (s Service) Run() error {
	return send(s.Body)
}

func send(body message.Body) error {
	m := message.GetType()
	if m == nil || !message.Enabled() {
		return errors.New("config message type error or not enabled")
	}
	if err := m.Send(body); err != nil {
		return err
	}
	return nil
}
