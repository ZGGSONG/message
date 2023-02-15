package message

import (
	"log"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	conf, err := InitConfig()
	if err != nil {
		t.Fatalf("[test] failed to initialize config: %v", err)
	}
	GLO_CONF = conf

	var s = Service{Body: Body{
		Title:   "test title",
		Content: "this is content, time is " + time.Now().Format("2006-01-02 15:04:05"),
	}}
	if err = s.Run(); err != nil {
		t.Fatalf("[test] failed to send message: %v", err)
	}
	for {
		_conf := <-GLO_CONF_CH
		log.Printf("config changed: %v", _conf)
		GLO_CONF = _conf
	}
}
