package message

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Bark struct {
	url string
	key string
}

func (b *Bark) Send(body Body) error {
	//log.Printf("[bark] sending message...")
	var reqBody = BarkRequest{
		DeviceKey: b.key,
		Title:     body.Title,
		Body:      body.Content,
	}
	req, _ := json.Marshal(reqBody)
	resp, err := http.Post(b.url,
		"application/json; charset=utf-8",
		bytes.NewReader(req))
	if err != nil {
		//log.Fatalf("[bark] http post failed: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	//log.Printf("[bark] send successful")
	return nil
}
