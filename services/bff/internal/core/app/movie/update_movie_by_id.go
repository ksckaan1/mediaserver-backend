package movie

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/moviepb"
	"strings"
)

type UpdateMovieByID struct {
	movieClient moviepb.MovieServiceClient
}

func NewUpdateMovieByID(movieClient moviepb.MovieServiceClient) *UpdateMovieByID {
	return &UpdateMovieByID{
		movieClient: movieClient,
	}
}

type UpdateMovieByIDRequest struct {
	MovieID     string `params:"movie_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TmdbId      string `json:"tmdb_id"`
	MediaId     string `json:"media_id"`
}

func (h *UpdateMovieByID) Handle(ctx context.Context, req *UpdateMovieByIDRequest) (*any, int, error) {
	_, err := h.movieClient.UpdateMovieByID(ctx, &moviepb.UpdateMovieByIDRequest{
		MovieId:     req.MovieID,
		Title:       req.Title,
		Description: req.Description,
		TmdbId:      req.TmdbId,
		MediaId:     req.MediaId,
	})
	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			return nil, http.StatusNotFound, errors.New("movie not found")
		}
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("movieClient.UpdateMovieByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
