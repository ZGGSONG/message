package message

type Message interface {
	Send(body Body) error
}

type Body struct {
	Title   string
	Content string
}

func GetType() Message {
	switch GLO_CONF.MsgType {
	case "bark":
		return &Bark{
			url: GLO_CONF.BarkUrl,
			key: GLO_CONF.BarkKey,
		}
	case "mail":
		return &Mail{
			Host:     GLO_CONF.MailHost,
			Port:     GLO_CONF.MailPort,
			Username: GLO_CONF.MailUser,
			Password: GLO_CONF.MailPwd,
			FromName: GLO_CONF.MailFromName,
			To:       GLO_CONF.MailTo,
		}
	}
	return nil
}

func Enabled() bool {
	return GLO_CONF.MsgEnabled
}
