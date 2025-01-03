package model

import "time"

type User struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
}

type UserList struct {
	Users  []User `json:"users"`
	Count  int64  `json:"count"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}
