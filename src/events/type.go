package event

import (
	"time"
) //import

type Event_t struct {
	Name      string    `json:"event_name"`
	CreatedAt time.Time `json:"created_at"`
	UserId    uint      `json:"user_id"`
	Email     string    `json:"email"`
	Metadata  string    `json:"metadata"`
} //Event_t
