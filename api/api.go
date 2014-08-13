package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Origin = "https://getpocket.com"

// Client represents a Pocket client that grants OAuth access to your application
type Client struct {
	ConsumerKey string
	AccessToken string
}

type AuthInfo struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

func NewClient(consumerKey string, accessToken string) *Client {
	return &Client{
		ConsumerKey: consumerKey,
		AccessToken: accessToken,
	}
}

func RequestJSON(action string, params interface{}, v interface{}) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	log.Println(string(body))

	req, err := http.NewRequest("POST", Origin+action, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
