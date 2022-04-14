package model

import "time"

const (
	AuthorizeSuccess = iota
	AuthorizeWrongPassword
	AuthorizeBlockedUser
)

type AuthorizationEvent struct {
	Timestamp time.Time `json: "id"`
	UserID    int       `json: "user_id"`
	Event     uint8     `json: "event"`
}