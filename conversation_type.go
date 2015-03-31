package intercom

type MessageType_t struct {
	MessageType string `json:"type,omitempty"`
	Id          string `json:"id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Email       string `json:"email,omitempty"`
} //MessageType_t

type InitiateConversation_t struct {
	MessageType string        `json"message_type,omitempty"`
	Subject     string        `json:"subject,omitempty"`
	Body        string        `json:"body,omitempty"`
	Template    string        `json:"template,omitempty"`
	From        MessageType_t `json:"from,omitempty"`
	To          MessageType_t `json:"to,omitempty"`
} //InitiateConversation_t
