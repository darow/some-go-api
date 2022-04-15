package model

import "time"

const (
	AuthorizeSuccess = iota
	AuthorizeWrongPassword
	AuthorizeBlockedUser
)

type AuthorizationLog struct {
	Timestamp time.Time `json: "id"`
	UserID    int       `json: "user_id,omitempty"`
	Event     uint8     `json: "event"`
}
