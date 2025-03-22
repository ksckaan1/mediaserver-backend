package mongodb

import (
	"context"
	"fmt"
	"series_service/internal/domain/core/customerrors"
	"series_service/internal/domain/core/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	seriesColl       *mongo.Collection
	seriesSearchColl *mongo.Collection
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		seriesColl:       db.Collection("series"),
		seriesSearchColl: db.Collection("series_search"),
	}
}

func (r *Repository) CreateSeries(ctx context.Context, series *models.Series) error {
	series.CreatedAt = time.Now()
	series.UpdatedAt = time.Now()
	_, err := r.seriesColl.InsertOne(ctx, series)
	if err != nil {
		return fmt.Errorf("seriesColl.InsertOne: %w", err)
	}
	_, err = r.seriesSearchColl.InsertOne(ctx, &models.SeriesSearch{
		ID:    series.ID,
		Title: series.Title,
	})
	if err != nil {
		return fmt.Errorf("seriesSearchColl.InsertOne: %w", err)
	}
	return nil
}

func (r *Repository) GetSeriesByID(ctx context.Context, id string) (*models.Series, error) {
	var series models.Series
	err := r.seriesColl.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&series)
	if err != nil {
		return nil, fmt.Errorf("seriesColl.FindOne: %w", err)
	}
	return &series, nil
}

func (r *Repository) ListSeries(ctx context.Context, limit int64, offset int64) (*models.SeriesList, error) {
	filter := bson.D{}
	count, err := r.seriesColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("seriesColl.CountDocuments: %w", err)
	}
	if count == 0 {
		return &models.SeriesList{
			List:   []*models.Series{},
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	opts := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := r.seriesColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("seriesColl.Find: %w", err)
	}
	var seriesList []*models.Series
	err = cursor.All(ctx, &seriesList)
	if err != nil {
		return nil, fmt.Errorf("cursor.All: %w", err)
	}
	return &models.SeriesList{
		List:   seriesList,
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (r *Repository) ListSeriesWithIDs(ctx context.Context, ids []string, limit int64, offset int64) (*models.SeriesList, error) {
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}
	count, err := r.seriesColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("seriesColl.CountDocuments: %w", err)
	}
	if count == 0 {
		return &models.SeriesList{
			List:   []*models.Series{},
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	opts := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := r.seriesColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("seriesColl.Find: %w", err)
	}
	var seriesList []*models.Series
	err = cursor.All(ctx, &seriesList)
	if err != nil {
		return nil, fmt.Errorf("cursor.All: %w", err)
	}
	return &models.SeriesList{
		List:   seriesList,
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (r *Repository) ListSeriesSearch(ctx context.Context) ([]*models.SeriesSearch, error) {
	cursor, err := r.seriesSearchColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("seriesSearchColl.Find: %w", err)
	}
	var seriesSearchList []*models.SeriesSearch
	err = cursor.All(ctx, &seriesSearchList)
	if err != nil {
		return nil, fmt.Errorf("cursor.All: %w", err)
	}
	return seriesSearchList, nil
}

func (r *Repository) UpdateSeriesByID(ctx context.Context, series *models.Series) error {
	series.UpdatedAt = time.Now()
	result, err := r.seriesColl.UpdateOne(ctx, bson.D{{Key: "_id", Value: series.ID}}, bson.D{{Key: "$set", Value: series}})
	if err != nil {
		return fmt.Errorf("seriesColl.UpdateOne: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("seriesColl.UpdateOne: %w", customerrors.ErrRecordNotFound)
	}
	_, err = r.seriesSearchColl.UpdateOne(ctx, bson.D{{Key: "_id", Value: series.ID}}, bson.D{{Key: "$set", Value: &models.SeriesSearch{
		ID:    series.ID,
		Title: series.Title,
	}}}, options.UpdateOne().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("seriesSearchColl.UpdateOne: %w", err)
	}
	return nil
}

func (r *Repository) DeleteSeriesByID(ctx context.Context, id string) error {
	result, err := r.seriesColl.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return fmt.Errorf("seriesColl.DeleteOne: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("seriesColl.DeleteOne: %w", customerrors.ErrRecordNotFound)
	}
	_, err = r.seriesSearchColl.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return fmt.Errorf("seriesSearchColl.DeleteOne: %w", err)
	}
	return nil
}
