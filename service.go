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
	m, err := GetType()
	if err != nil {
		return err
	}
	if !Enabled() {
		return errors.New("config set message not enabled")
	}
	if err := m.Send(body); err != nil {
		return err
	}
	return nil
}
