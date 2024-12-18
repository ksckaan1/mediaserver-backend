package model

import (
	"mediaserver/internal/domain/core/enum/mediatype"
	"mediaserver/internal/domain/core/enum/storagetype"
	"time"
)

type Media struct {
	ID          string                  `json:"id"`
	CreatedAt   time.Time               `json:"created_at"`
	Path        string                  `json:"path"`
	Type        mediatype.MediaType     `json:"type"`
	StorageType storagetype.StorageType `json:"storage_type"`
	MimeType    string                  `json:"mime_type"`
	Size        int64                   `json:"size"`
}

type MediaList struct {
	List   []*Media `json:"list"`
	Count  int64    `json:"count"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
}

type FileInfo struct {
	ID          string                  `json:"id"`
	Path        string                  `json:"path"`
	Type        mediatype.MediaType     `json:"type"`
	StorageType storagetype.StorageType `json:"storage_type"`
	MimeType    string                  `json:"mime_type"`
	Size        int64                   `json:"size"`
}
