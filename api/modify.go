package api

type Action struct {
	Action string `json:"action"`
	ItemId int    `json:"item_id,string"`
}

type ModifyAPIResponse struct {
	ActionResults []bool
	Status        int
}

type ModifyAPIOptionsWithAuth struct {
	Actions []*Action `json:"actions"`
	AuthInfo
}

func NewArchiveAction(itemId int) *Action {
	return &Action{
		Action: "archive",
		ItemId: itemId,
	}
}

func (c *Client) Modify(actions ...*Action) (*ModifyAPIResponse, error) {
	res := &ModifyAPIResponse{}
	data := ModifyAPIOptionsWithAuth{
		AuthInfo: c.authInfo(),
		Actions:  actions,
	}
	err := PostJSON("/v3/send", data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
