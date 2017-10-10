package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Origin is the constant origin URL for the Pocket API
var Origin = "https://getpocket.com"

// DefaultClient is the client used for making all requests
var DefaultClient = http.DefaultClient

// Client represents a Pocket client that grants OAuth access to your application
type Client struct {
	authInfo
}

type authInfo struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

// NewClient creates a new Pocket client.
func NewClient(consumerKey, accessToken string) *Client {
	return &Client{
		authInfo: authInfo{
			ConsumerKey: consumerKey,
			AccessToken: accessToken,
		},
	}
}

func doJSON(req *http.Request, res interface{}) error {
	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(res)
}

// PostJSON posts the data to the API endpoint, storing the result in res.
func PostJSON(action string, data, res interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", Origin+action, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return doJSON(req, res)
}
