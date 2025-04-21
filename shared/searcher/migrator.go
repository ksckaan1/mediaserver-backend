package searcher

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/typesense/typesense-go/v3/typesense/api"
	"github.com/typesense/typesense-go/v3/typesense/api/pointer"
)

func (s *Searcher) migrateMoviesCollection(ctx context.Context) error {
	fields := []api.Field{
		{
			Index: pointer.True(),
			Name:  "id",
			Sort:  pointer.True(),
			Type:  "string",
		},
		{
			Index: pointer.True(),
			Name:  "title",
			Sort:  pointer.True(),
			Type:  "string",
		},
		{
			Index: pointer.True(),
			Name:  "tags",
			Type:  "string[]",
		},
	}
	collections, err := s.client.Collections().Retrieve(ctx)
	if err != nil {
		return fmt.Errorf("Collections().Retrieve: %w", err)
	}
	_, ok := lo.Find(collections, func(collection *api.CollectionResponse) bool {
		return collection.Name == "movies"
	})
	if !ok {
		_, err := s.client.Collections().Create(ctx, &api.CollectionSchema{
			Name:               "movies",
			EnableNestedFields: pointer.True(),
			Fields:             fields,
		})
		if err != nil {
			return fmt.Errorf("Collections().Create: %w", err)
		}
	} else {
		s.client.Collection("movies").Update(ctx, &api.CollectionUpdateSchema{
			Fields: fields,
		})
	}
	return nil
}

func (s *Searcher) migrateSeriesCollection(ctx context.Context) error {
	fields := []api.Field{
		{
			Index: pointer.True(),
			Name:  "id",
			Sort:  pointer.True(),
			Type:  "string",
		},
		{
			Index: pointer.True(),
			Name:  "title",
			Sort:  pointer.True(),
			Type:  "string",
		},
		{
			Index: pointer.True(),
			Name:  "tags",
			Type:  "string[]",
		},
	}
	collections, err := s.client.Collections().Retrieve(ctx)
	if err != nil {
		return fmt.Errorf("Collections().Retrieve: %w", err)
	}
	_, ok := lo.Find(collections, func(collection *api.CollectionResponse) bool {
		return collection.Name == "series"
	})
	if !ok {
		_, err := s.client.Collections().Create(ctx, &api.CollectionSchema{
			Name:               "series",
			EnableNestedFields: pointer.True(),
			Fields:             fields,
		})
		if err != nil {
			return fmt.Errorf("Collections().Create: %w", err)
		}
	} else {
		s.client.Collection("movies").Update(ctx, &api.CollectionUpdateSchema{
			Fields: fields,
		})
	}
	return nil
}
