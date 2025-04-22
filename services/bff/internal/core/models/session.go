package models

import (
	"shared/enums/usertype"
	"time"
)

type Session struct {
	SessionId string
	UserId    string
	UserAgent string
	UserType  usertype.UserType
	IpAddress string
	CreatedAt time.Time
}
