package mongodb

import (
	"context"
	"fmt"
	"movie_service/internal/domain/core/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	moviesColl       *mongo.Collection
	moviesSearchColl *mongo.Collection
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		moviesColl:       db.Collection("movies"),
		moviesSearchColl: db.Collection("movies_search"),
	}
}

func (r *Repository) CreateMovie(ctx context.Context, movie *models.Movie) error {
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()
	_, err := r.moviesColl.InsertOne(ctx, movie)
	if err != nil {
		return fmt.Errorf("moviesColl.InsertOne: %w", err)
	}
	_, err = r.moviesSearchColl.InsertOne(ctx, &models.MovieSearch{
		ID:    movie.ID,
		Title: movie.Title,
	})
	if err != nil {
		return fmt.Errorf("moviesSearchColl.InsertOne: %w", err)
	}
	return nil
}

func (r *Repository) GetMovieByID(ctx context.Context, id string) (*models.Movie, error) {
	result := r.moviesColl.FindOne(ctx, bson.M{"_id": id})
	if result.Err() != nil {
		return nil, fmt.Errorf("moviesColl.FindOne: %w", result.Err())
	}
	var movie models.Movie
	if err := result.Decode(&movie); err != nil {
		return nil, fmt.Errorf("result.Decode: %w", err)
	}
	return &movie, nil
}

func (r *Repository) ListMovies(ctx context.Context, limit, offset int64) (*models.MovieList, error) {
	filter := bson.M{}
	count, err := r.moviesColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("moviesColl.CountDocuments: %w", err)
	}
	opt := options.Find().
		SetLimit(limit).
		SetSkip(offset)
	cursor, err := r.moviesColl.Find(ctx, filter, opt)
	if err != nil {
		return nil, fmt.Errorf("moviesColl.Find: %w", err)
	}
	var movies []*models.Movie
	err = cursor.All(ctx, &movies)
	if err != nil {
		return nil, fmt.Errorf("cursor.All: %w", err)
	}
	return &models.MovieList{
		List:   movies,
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (r *Repository) ListMoviesWithIDs(ctx context.Context, ids []string, limit, offset int64) (*models.MovieList, error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	count, err := r.moviesColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("moviesColl.CountDocuments: %w", err)
	}
	if count == 0 {
		return &models.MovieList{
			List:   []*models.Movie{},
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	var movies []*models.Movie
	opts := options.Find().SetLimit(limit).SetSkip(offset)
	result, err := r.moviesColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("moviesColl.Find: %w", err)
	}
	err = result.All(ctx, &movies)
	if err != nil {
		return nil, fmt.Errorf("result.All: %w", err)
	}
	return &models.MovieList{
		List:   movies,
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (r *Repository) ListMoviesSearch(ctx context.Context) ([]*models.MovieSearch, error) {
	cursor, err := r.moviesSearchColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("moviesSearchColl.Find: %w", err)
	}
	var moviesSearch []*models.MovieSearch
	err = cursor.All(ctx, &moviesSearch)
	if err != nil {
		return nil, fmt.Errorf("cursor.All: %w", err)
	}
	return moviesSearch, nil
}

func (r *Repository) UpdateMovieByID(ctx context.Context, movie *models.Movie) error {
	movie.UpdatedAt = time.Now()
	result, err := r.moviesColl.UpdateOne(ctx, bson.M{"_id": movie.ID}, bson.M{"$set": movie})
	if err != nil {
		return fmt.Errorf("moviesColl.UpdateOne: %w", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("movie not found")
	}
	_, err = r.moviesSearchColl.UpdateOne(ctx, bson.M{"_id": movie.ID}, bson.M{"$set": &models.MovieSearch{
		ID:    movie.ID,
		Title: movie.Title,
	}}, options.UpdateOne().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("moviesSearchColl.UpdateOne: %w", err)
	}
	return nil
}

func (r *Repository) DeleteMovieByID(ctx context.Context, id string) error {
	result, err := r.moviesColl.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("moviesColl.DeleteOne: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("movie not found")
	}
	_, err = r.moviesSearchColl.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("moviesSearchColl.DeleteOne: %w", err)
	}
	return nil
}
