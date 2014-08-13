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

type State string

const (
	StateUnread  State = "unread"
	StateArchive       = "archive"
	StateAll           = "all"
)

type ContentType string

const (
	ContentTypeArticle ContentType = "article"
	ContentTypeVideo               = "video"
	ContentTypeImage               = "image"
)

type Sort string

const (
	SortNewest Sort = "newest"
	SortOldest      = "oldest"
	SortTitle       = "title"
	SortSite        = "site"
)

type DetailType string

const (
	DetailTypeSimple   DetailType = "simple"
	DetailTypeComplete            = "complete"
)

type FavoriteFilter string

const (
	FavoriteFilterUnspecified FavoriteFilter = ""
	FavoriteFilterUnfavorited                = "0"
	FavoriteFilterFavorited                  = "1"
)

type RetrieveAPIOption struct {
	State       State          `json:"state,omitempty"`
	Favorite    FavoriteFilter `json:"favorite,omitempty"`
	Tag         string         `json:"tag,omitempty"`
	ContentType ContentType    `json:"contentType,omitempty"`
	Sort        Sort           `json:"sort,omitempty"`
	DetailType  DetailType     `json:"detailType,omitempty"`
	Search      string         `json:"search,omitempty"`
	Domain      string         `json:"domain,omitempty"`
	Since       int            `json:"since,omitempty"`
	Count       int            `json:"count,omitempty"`
	Offset      int            `json:"offset,omitempty"`
}

type AuthInfo struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

type RetrieveAPIOptionWithAuth struct {
	*RetrieveAPIOption
	AuthInfo
}

// URL is an alias for ResolvedURL
func (item *Item) URL() string {
	return item.ResolvedURL
}

// Title is an alias for ResolvedTitle
func (item *Item) Title() string {
	return item.ResolvedTitle
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

func (c *Client) Retrieve(options *RetrieveAPIOption) (*RetrieveAPIResponse, error) {
	params := RetrieveAPIOptionWithAuth{
		AuthInfo: AuthInfo{
			ConsumerKey: c.ConsumerKey,
			AccessToken: c.AccessToken,
		},
		RetrieveAPIOption: options,
	}

	res := &RetrieveAPIResponse{}
	err := RequestJSON("/v3/get", params, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
