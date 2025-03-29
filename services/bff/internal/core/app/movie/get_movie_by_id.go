package movie

import (
	"bff-service/internal/core/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/enums/mediatype"
	"shared/pb/moviepb"
	"strings"
)

type GetMovieByID struct {
	movieClient moviepb.MovieServiceClient
}

func NewGetMovieByID(movieClient moviepb.MovieServiceClient) *GetMovieByID {
	return &GetMovieByID{
		movieClient: movieClient,
	}
}

type GetMovieByIDRequest struct {
	MovieID string `params:"movie_id"`
}

type GetMovieByIDResponse struct {
	ID          string        `json:"id"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	MediaInfo   *models.Media `json:"media_info"`
	TmdbInfo    *models.TMDB  `json:"tmdb_info"`
}

func (h *GetMovieByID) Handle(ctx context.Context, req *GetMovieByIDRequest) (*models.Movie, int, error) {
	resp, err := h.movieClient.GetMovieByID(ctx, &moviepb.GetMovieByIDRequest{
		MovieId: req.MovieID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			return nil, http.StatusNotFound, errors.New("movie not found")
		}
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("movieClient.GetMovieByID: %w", err)
	}

	var mediaInfo *models.Media
	if resp.MediaInfo != nil {
		mediaInfo = &models.Media{
			ID:        resp.MediaInfo.Id,
			CreatedAt: resp.MediaInfo.CreatedAt.AsTime(),
			UpdatedAt: resp.MediaInfo.UpdatedAt.AsTime(),
			Title:     resp.MediaInfo.Title,
			Path:      resp.MediaInfo.Path,
			MimeType:  resp.MediaInfo.MimeType,
			Size:      resp.MediaInfo.Size,
			Type:      mediatype.FromNumber(int32(resp.MediaInfo.Type)),
		}
	}

	var tmdbInfo *models.TMDB
	if resp.TmdbInfo != nil {
		tmdbInfo = &models.TMDB{
			Id:   resp.TmdbInfo.Id,
			Data: resp.TmdbInfo.Data.AsMap(),
		}
	}

	return &models.Movie{
		ID:          resp.Id,
		CreatedAt:   resp.CreatedAt.AsTime(),
		UpdatedAt:   resp.UpdatedAt.AsTime(),
		Title:       resp.Title,
		Description: resp.Description,
		MediaInfo:   mediaInfo,
		TmdbInfo:    tmdbInfo,
	}, http.StatusOK, nil
}
