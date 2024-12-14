package movie

import (
	"context"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/enum/mediatype"
)

func (m *Movie) validateMedia(ctx context.Context, mediaID string) error {
	if mediaID != "" {
		media, err := m.repo.GetMediaByID(ctx, mediaID)
		if err != nil {
			return fmt.Errorf("repo.GetMediaByID: %w", err)
		}
		if media.Type != mediatype.Video {
			return fmt.Errorf("check media type: %w", customerrors.ErrInvalidMediaType)
		}
	}
	return nil
}
