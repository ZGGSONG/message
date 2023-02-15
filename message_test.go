package message

import (
	"log"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	conf, err := InitConfig()
	if err != nil {
		t.Fatalf("[test] failed to initialize config: %v", err)
	}
	GLO_CONF = conf

	m, err := GetType()
	if !Enabled() || err != nil {
		t.Fatalf("[test] failed to get type %v", err)
	}
	if err = m.Send(Body{
		Title:   "Hello",
		Content: "World " + time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		log.Fatalf("[test] failed to send fail: %v", err)
	}
}
