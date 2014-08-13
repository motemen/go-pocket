package api

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

	// Fields for detailed response
	Tags    map[string]map[string]interface{}
	Authors map[string]map[string]interface{}
	Images  map[string]map[string]interface{}
	Videos  map[string]map[string]interface{}
}
