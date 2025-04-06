package models

import "time"

type Session struct {
	SessionID string    `json:"session_id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id"`
	UserAgent string    `json:"user_agent"`
	IpAddress string    `json:"ip_address"`
}
