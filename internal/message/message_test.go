package message

import (
	"github.com/zggsong/message/internal/config"
	"github.com/zggsong/message/internal/global"
	"log"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	conf, err := config.InitConfig()
	if err != nil {
		log.Println("[test] failed to initialize config: ", err)
	}
	global.GLO_CONF = conf

	m := GetType()
	if !Enabled() || m == nil {
		t.Fatalf("[test] failed to get type...")
	}
	if err = m.Send(Body{
		Title:   "Hello",
		Content: "World " + time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		log.Fatalf("[test] failed to send fail: %v", err)
	}
}
