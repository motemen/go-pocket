package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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

type RetrieveAPIOption struct {
	State       State
	Favorite    bool
	Tag         string
	ContentType ContentType
	Sort        Sort
	DetailType  DetailType
	Search      string
	Domain      string
	Since       time.Time
	Count       int
	Offset      int
}

type RetrieveAPIResponse struct {
	Status int
	List   map[string]*Item
}

type ItemStatus int

const (
	ItemStatusUnread   ItemStatus = 0
	ItemStatusArchived            = 1
	ItemStatusDeleted             = 2
)

type ItemMediaAttachment int

const (
	ItemMediaAttachmentNoMedia  ItemMediaAttachment = 0
	ItemMediaAttachmentHasMedia                     = 1
	ItemMediaAttachmentIsMedia                      = 2
)

type Item struct {
	ItemId        int        `json:"item_id,string"`
	ResolvedId    int        `json:"resolved_id,string"`
	GivenURL      string     `json:"given_url"`
	ResolvedURL   string     `json:"resolved_url"`
	GivenTitle    string     `json:"given_title"`
	ResolvedTitle string     `json:"resolved_title"`
	Favorite      int        `json:",string"`
	Status        ItemStatus `json:",string"`
	Excerpt       string
	IsArticle     int                 `json:"is_article,string"`
	HasImage      ItemMediaAttachment `json:"has_image,string"`
	HasVideo      ItemMediaAttachment `json:"has_video,string"`
	WordCount     int                 `json:"word_count,string"`
	Tags          []interface{}
	Authors       []interface{}
	Images        []interface{}
	Videos        []interface{}
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

func requestRaw(action string, params url.Values) (io.Reader, error) {
	req, err := http.NewRequest("POST", Origin+action, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	return resp.Body, nil
}

func Request(action string, params url.Values, v interface{}) error {
	r, err := requestRaw(action, params)
	if err != nil {
		return err
	}

	d := json.NewDecoder(r)
	return d.Decode(v)
}

func (c *Client) Retrieve(options *RetrieveAPIOption) (*RetrieveAPIResponse, error) {
	params := url.Values{
		"consumer_key": {c.ConsumerKey},
		"access_token": {c.AccessToken},
	}

	if options.Domain != "" {
		params.Add("domain", options.Domain)
	}

	if options.Search != "" {
		params.Add("search", options.Search)
	}

	res := &RetrieveAPIResponse{}
	err := Request(
		"/v3/get",
		params,
		res,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
