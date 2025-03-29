package movie

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/moviepb"
	"strings"
)

type CreateMovie struct {
	movieClient moviepb.MovieServiceClient
}

func NewCreateMovie(movieClient moviepb.MovieServiceClient) *CreateMovie {
	return &CreateMovie{
		movieClient: movieClient,
	}
}

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	TmdbId      string `json:"tmdb_id"`
	MediaId     string `json:"media_id"`
}

type CreateMovieResponse struct {
	MovieID string `json:"movie_id"`
}

func (h *CreateMovie) Handle(ctx context.Context, req *CreateMovieRequest) (*CreateMovieResponse, int, error) {
	resp, err := h.movieClient.CreateMovie(ctx, &moviepb.CreateMovieRequest{
		Title:       req.Title,
		Description: req.Description,
		TmdbId:      req.TmdbId,
		MediaId:     req.MediaId,
	})
	if err != nil {
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("movieClient.CreateMovie: %w", err)
	}
	return &CreateMovieResponse{
		MovieID: resp.MovieId,
	}, http.StatusCreated, nil
}
