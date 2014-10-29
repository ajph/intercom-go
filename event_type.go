package event

import (
	"time"
) //import

type Event_t struct {
	EventName string    `json:"event_name"`
	CreatedAt time.Time `json:"created_at"`
	UserId    uint      `json:"user_id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Metadata  string    `json:"metadata,omitempty"`
} //Event_t
