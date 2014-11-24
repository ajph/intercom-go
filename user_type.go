package intercom

type User_t struct {
	Type                   string                 `json:"type"`
	Id                     string                 `json:"id"`
	CreatedAt              int64                  `json:"created_at"`
	RemoteCreatedAt        int64                  `json:"remote_created_at,omitempty"`
	UpdatedAt              int64                  `json:"updated_at,omitempty"`
	UserId                 string                 `json:"user_id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
	LastRequestAt          int64                  `json:"last_request_at,omitempty"`
	SessionCount           int64                  `json:"session_count,omitempty"`
	Avatar                 Avator_t               `json:"avatar,omitempty"`
	UnsubscribedFromEmails bool                   `json:"unsubscribed_from_emails,omitempty"`
} //User_t

type Avator_t struct {
} //Avator_t
