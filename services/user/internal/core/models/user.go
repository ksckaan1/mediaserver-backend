package models

import "time"

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}

type UserList struct {
	List   []*User `json:"list"`
	Count  int64   `json:"count"`
	Limit  int64   `json:"limit"`
	Offset int64   `json:"offset"`
}
