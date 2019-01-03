package liveagent

type Conversation struct {
	Message       string `url:"message"`
	UserID        string `url:"useridentifier"`
	DeptID        string `url:"department"`
	Subject       string `url:"subject"`
	Recipient     string `url:"recipient"`
	DoNotSendMail string `url:"do_not_send_mail"`
	UseTemplate   string `url:"use_template"`
	IsHTMLMessage string `url:"is_html_message"`
	Status        string `url:"status"`
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
	var r NewConversationResponse
	err := c.post("conversations", conv, &r)
	if err != nil {
		return nil, err
	}
	return &(r.Response), nil
}
