package message

import (
	"errors"
)

type Service struct {
	Body Body
}

func (s Service) Run() error {
	return send(s.Body)
}

func send(body Body) error {
	m := GetType()
	if m == nil || !Enabled() {
		return errors.New("config message type error or not enabled")
	}
	if err := m.Send(body); err != nil {
		return err
	}
	return nil
}
