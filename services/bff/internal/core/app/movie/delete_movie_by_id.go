package movie

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/moviepb"
	"strings"
)

type DeleteMovieByID struct {
	movieClient moviepb.MovieServiceClient
}

func NewDeleteMovieByID(movieClient moviepb.MovieServiceClient) *DeleteMovieByID {
	return &DeleteMovieByID{
		movieClient: movieClient,
	}
}

type DeleteMovieByIDRequest struct {
	MovieID string `params:"movie_id"`
}

func (h *DeleteMovieByID) Handle(ctx context.Context, req *DeleteMovieByIDRequest) (*any, int, error) {
	_, err := h.movieClient.DeleteMovieByID(ctx, &moviepb.DeleteMovieByIDRequest{
		MovieId: req.MovieID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			return nil, http.StatusNotFound, errors.New("movie not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("movieClient.DeleteMovieByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
