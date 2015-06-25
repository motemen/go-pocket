package api

// AddOption is the options for the Add API.
type AddOption struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
	Tags  string `json:"tags,omitempty"`
}

type addAPIOptionWithAuth struct {
	*AddOption
	authInfo
}

// AddResult is a struct with no fields, which is passed in below to conform
// to the interface for PostJSON.
type AddResult struct{}

// Add only returns an error status, since adding an article doesn't have
// any other meaningful return value.
func (c *Client) Add(options *AddOption) error {
	data := addAPIOptionWithAuth{
		authInfo:  c.authInfo,
		AddOption: options,
	}

	res := &AddResult{}
	err := PostJSON("/v3/add", data, res)
	if err != nil {
		return nil
	}

	return err
}
