package config

import (
	"github.com/zggsong/message/internal/global"
	"github.com/zggsong/message/internal/model"
	"testing"
)

func TestConfig(t *testing.T) {
	global.GLO_CONF_CH = make(chan model.Config)
	config, err := InitConfig()
	if err != nil {
		t.Fatalf("Failed to initialize config: %v", err)
	}
	t.Log(config)

	var count int
	for {
		if count > 5 {
			return
		}
		config = <-global.GLO_CONF_CH
		count++
		t.Log(config)
	}
}
