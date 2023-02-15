package message

import (
	"github.com/zggsong/message/internal/config"
	"github.com/zggsong/message/internal/global"
	"github.com/zggsong/message/internal/message"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	conf, err := config.InitConfig()
	if err != nil {
		t.Fatalf("[test] failed to initialize config: %v", err)
	}
	global.GLO_CONF = conf

	var s = Service{Body: message.Body{
		Title:   "test title",
		Content: "this is content, time is " + time.Now().Format("2006-01-02 15:04:05"),
	}}
	if err = s.Run(); err != nil {
		t.Fatalf("[test] failed to send message: %v", err)
	}
}
