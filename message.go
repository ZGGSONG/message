package message

import "errors"

type Message interface {
	Send(body Body) error
}

type Body struct {
	Title   string
	Content string
}

func GetType() (Message, error) {
	switch GLO_CONF.MsgType {
	case "bark":
		if GLO_CONF.BarkUrl == "" || GLO_CONF.BarkKey == "" {
			return nil, errors.New("bark conf not completed")
		}
		return &Bark{
			url: GLO_CONF.BarkUrl,
			key: GLO_CONF.BarkKey,
		}, nil
	case "mail":
		if GLO_CONF.MailHost == "" || GLO_CONF.MailPort == 0 || GLO_CONF.MailUser == "" || GLO_CONF.MailPwd == "" || GLO_CONF.MailFromName == "" || GLO_CONF.MailTo == nil {
			return nil, errors.New("mail conf not completed")
		}
		return &Mail{
			Host:     GLO_CONF.MailHost,
			Port:     GLO_CONF.MailPort,
			Username: GLO_CONF.MailUser,
			Password: GLO_CONF.MailPwd,
			FromName: GLO_CONF.MailFromName,
			To:       GLO_CONF.MailTo,
		}, nil
	}
	return nil, errors.New("message type in config is not supported")
}

func Enabled() bool {
	return GLO_CONF.MsgEnabled
}
