package models

import "time"

type Session struct {
	SessionId string
	UserId    string
	UserAgent string
	IpAddress string
	CreatedAt time.Time
}
