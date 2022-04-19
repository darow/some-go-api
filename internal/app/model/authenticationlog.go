package model

import (
	"encoding/json"
	"time"
)

const (
	AuthorizeSuccess = iota
	AuthorizeWrongPassword
	AuthorizeBlockedUser
)

type AuthenticationLog struct {
	Timestamp time.Time `json:"timestamp"`
	UserID    int       `json:"-"`
	Event     uint8     `json:"event"`
}

func (d *AuthenticationLog) MarshalJSON() ([]byte, error) {
	events := map[uint8]string{
		0: "AuthorizeSuccess",
		1: "AuthorizeWrongPassword",
		2: "AuthorizeBlockedUser",
	}
	type Alias AuthenticationLog
	return json.Marshal(&struct {
		*Alias
		Timestamp string `json:"timestamp"`
		Event     string `json:"event"`
	}{
		Alias:     (*Alias)(d),
		Timestamp: d.Timestamp.Format("2006/01/02 15:04:05"),
		Event:     events[d.Event],
	})
}
