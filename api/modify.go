package api

// Action represents one action in a bulk modify requests.
type Action struct {
	Action string `json:"action"`
	ItemID int    `json:"item_id,string"`
}

// NewArchiveAction creates an acrhive action.
func NewArchiveAction(itemID int) *Action {
	return &Action{
		Action: "archive",
		ItemID: itemID,
	}
}

// ModifyResult represents the modify API's result.
type ModifyResult struct {
	// The results for each of the requested actions.
	ActionResults []bool
	Status        int
}

type modifyAPIOptionsWithAuth struct {
	Actions []*Action `json:"actions"`
	authInfo
}

// Modify requests bulk modification on items.
func (c *Client) Modify(actions ...*Action) (*ModifyResult, error) {
	res := &ModifyResult{}
	data := modifyAPIOptionsWithAuth{
		authInfo: c.authInfo,
		Actions:  actions,
	}
	err := PostJSON("/v3/send", data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
