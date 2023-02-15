package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Bark struct {
	url string
	key string
}

func (b *Bark) Send(body Body) error {
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
		return errors.New(fmt.Sprintf("bark http post failed: %v", err))
	}
	defer resp.Body.Close()

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("bark read resp body failed, %v", err))
	}

	var response BarkResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return errors.New(fmt.Sprintf("bark json unmarshal failed, %v", err))
	}

	if response.Code != 200 {
		return errors.New("bar send failed")
	}
	//log.Printf("[bark] send successful")
	return nil
}
