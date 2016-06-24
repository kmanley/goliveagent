package liveagent

type Conversation struct {
	Message   string `url:"message"`
	UserID    string `url:"useridentifier"`
	DeptID    string `url:"department"`
	Subject   string `url:"subject"`
	Recipient string `url:"recipient"`
	APIKey    string `url:"apikey"`
}

type NewConversation struct {
	Status         string `json:"status"`
	StatusCode     int    `json:"statuscode"`
	ConversationID string `json:"conversationid"`
	Code           string `json:"code"`
	PublicURLCode  string `json:"publicurlcode"`
}

type NewConversationResponse struct {
	Response NewConversation `json:"response"`
}

func (c *Client) ConversationCreate(conv *Conversation) (*NewConversation, error) {
	conv.APIKey = c.APIKey
	var r NewConversationResponse
	err := c.post("conversations", conv, &r)
	if err != nil {
		return nil, err
	}
	return &(r.Response), nil
}
