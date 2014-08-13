package api

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

type RetrieveAPIOptionWithAuth struct {
	*RetrieveAPIOption
	AuthInfo
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
