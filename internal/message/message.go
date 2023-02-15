package message

import "github.com/zggsong/message/internal/global"

type Message interface {
	Send(body Body) error
}

type Body struct {
	Title   string
	Content string
}

func GetType() Message {
	switch global.GLO_CONF.MsgType {
	case "bark":
		return &Bark{
			url: global.GLO_CONF.BarkUrl,
			key: global.GLO_CONF.BarkKey,
		}
	case "mail":
		return &Mail{
			Host:     global.GLO_CONF.MailHost,
			Port:     global.GLO_CONF.MailPort,
			Username: global.GLO_CONF.MailUser,
			Password: global.GLO_CONF.MailPwd,
			FromName: global.GLO_CONF.MailFromName,
			To:       global.GLO_CONF.MailTo,
		}
	}
	return nil
}

func Enabled() bool {
	return global.GLO_CONF.MsgEnabled
}
