package models

import (
	"shared/enums/usertype"
	"time"
)

type User struct {
	ID        string            `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Username  string            `json:"username"`
	Password  string            `json:"password,omitempty"`
	UserType  usertype.UserType `json:"user_type"`
}
