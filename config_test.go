package message

import (
	"testing"
)

func TestConfig(t *testing.T) {
	GLO_CONF_CH = make(chan Config)
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
		config = <-GLO_CONF_CH
		count++
		t.Log(config)
	}
}
