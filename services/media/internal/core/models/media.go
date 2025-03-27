package models

import (
	"shared/enums/mediatype"
	"time"
)

type Media struct {
	ID        string              `json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Title     string              `json:"title"`
	Path      string              `json:"path"`
	Type      mediatype.MediaType `json:"type"`
	MimeType  string              `json:"mime_type"`
	Size      int64               `json:"size"`
}

type MediaList struct {
	List   []*Media `json:"list"`
	Count  int64    `json:"count"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
}
