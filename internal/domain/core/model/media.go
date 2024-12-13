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
	Size        int64                   `json:"size"`
}
