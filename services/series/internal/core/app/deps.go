package app

import (
	"context"
	"series_service/internal/core/models"
)

type Repository interface {
	CreateSeries(ctx context.Context, series *models.Series) error
	GetSeriesByID(ctx context.Context, id string) (*models.Series, error)
	ListSeries(ctx context.Context, limit, offset int64) (*models.SeriesList, error)
	UpdateSeriesByID(ctx context.Context, series *models.Series) error
	DeleteSeriesByID(ctx context.Context, id string) error
}
type Searcher interface {
	AddDocument(ctx context.Context, collectionName string, doc any) error
	DeleteDocument(ctx context.Context, collectionName string, docID string) error
	UpdateDocument(ctx context.Context, collectionName string, id string, doc any) error
	Search(ctx context.Context, collectionName, query, queryBy string, limit, offset int, cache bool, v any) (int, error)
}
