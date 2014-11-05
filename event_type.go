package intercom

type Event_t struct {
	EventName string                 `json:"event_name"`
	CreatedAt int64                  `json:"created_at"`
	UserId    uint                   `json:"user_id,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
} //Event_t
