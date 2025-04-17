package movie

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/enums/mediatype"
	"shared/pb/moviepb"

	"github.com/samber/lo"
)

type ListMovies struct {
	movieClient moviepb.MovieServiceClient
}

func NewListMovies(movieClient moviepb.MovieServiceClient) *ListMovies {
	return &ListMovies{
		movieClient: movieClient,
	}
}

type ListMoviesRequest struct {
	Limit  *int64 `query:"limit"`
	Offset int64  `query:"offset"`
}

type ListMoviesResponse struct {
	List   []*models.Movie `json:"list"`
	Count  int64           `json:"count"`
	Limit  int64           `json:"limit"`
	Offset int64           `json:"offset"`
}

func (h *ListMovies) Handle(ctx context.Context, req *ListMoviesRequest) (*ListMoviesResponse, int, error) {
	var limit int64 = 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	resp, err := h.movieClient.ListMovies(ctx, &moviepb.ListMoviesRequest{
		Limit:  limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("movieClient.ListMovies: %w", err)
	}

	return &ListMoviesResponse{
		List: lo.Map(resp.List, func(movie *moviepb.Movie, _ int) *models.Movie {
			var mediaInfo *models.Media
			if movie.MediaInfo != nil {
				mediaInfo = &models.Media{
					ID:        movie.MediaInfo.Id,
					CreatedAt: movie.MediaInfo.CreatedAt.AsTime(),
					UpdatedAt: movie.MediaInfo.UpdatedAt.AsTime(),
					Title:     movie.MediaInfo.Title,
					Path:      movie.MediaInfo.Path,
					MimeType:  movie.MediaInfo.MimeType,
					Size:      movie.MediaInfo.Size,
					Type:      mediatype.FromNumber(int32(movie.MediaInfo.Type)),
				}
			}

			var tmdbInfo *models.TMDB
			if movie.TmdbInfo != nil {
				tmdbInfo = &models.TMDB{
					Id:   movie.TmdbInfo.Id,
					Data: movie.TmdbInfo.Data.AsMap(),
				}
			}

			return &models.Movie{
				ID:          movie.Id,
				CreatedAt:   movie.CreatedAt.AsTime(),
				UpdatedAt:   movie.UpdatedAt.AsTime(),
				Title:       movie.Title,
				Description: movie.Description,
				MediaInfo:   mediaInfo,
				TmdbInfo:    tmdbInfo,
				Tags:        movie.Tags,
			}
		}),
		Count:  resp.Count,
		Limit:  resp.Limit,
		Offset: resp.Offset,
	}, http.StatusOK, nil
}
