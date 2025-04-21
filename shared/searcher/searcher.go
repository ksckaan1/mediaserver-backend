package searcher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/typesense/typesense-go/v3/typesense"
	"github.com/typesense/typesense-go/v3/typesense/api"
	"github.com/typesense/typesense-go/v3/typesense/api/pointer"
)

type Searcher struct {
	client *typesense.Client
}

func New(typesenseURL, apiKey string) *Searcher {
	client := typesense.NewClient(
		typesense.WithServer(typesenseURL),
		typesense.WithAPIKey(apiKey))
	return &Searcher{
		client: client,
	}
}

func (s *Searcher) Migrate(ctx context.Context, collection string) error {
	switch collection {
	case "movies":
		return s.migrateMoviesCollection(ctx)
	case "series":
		return s.migrateSeriesCollection(ctx)
	default:
		return fmt.Errorf("unknown collection: %s", collection)
	}
}

func (s *Searcher) AddDocument(ctx context.Context, collectionName string, doc any) error {
	_, err := s.client.Collection(collectionName).Documents().Create(ctx, doc, &api.DocumentIndexParameters{})
	if err != nil {
		return fmt.Errorf("Collection(%s).Documents().Create: %w", collectionName, err)
	}
	return nil
}

func (s *Searcher) DeleteDocument(ctx context.Context, collectionName string, docID string) error {
	_, err := s.client.Collection(collectionName).Document(docID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("Collection(%s).Document(%s).Delete: %w", collectionName, docID, err)
	}
	return nil
}

func (s *Searcher) UpdateDocument(ctx context.Context, collectionName, id string, doc any) error {
	_, err := s.client.Collection(collectionName).Document(id).Update(ctx, doc, &api.DocumentIndexParameters{})
	if err != nil {
		return fmt.Errorf("Collection(%s).Document(%s).Update: %w", collectionName, id, err)
	}
	return nil
}

func (s *Searcher) Search(ctx context.Context, collectionName, query, queryBy string, limit, offset int, cache bool, v any) (int, error) {
	result, err := s.client.Collection(collectionName).Documents().Search(ctx, &api.SearchCollectionParams{
		Q:                &query,
		QueryBy:          &queryBy,
		Limit:            &limit,
		Offset:           &offset,
		ExhaustiveSearch: pointer.True(),
		UseCache:         &cache,
		ExcludeFields:    pointer.String(""),
	})
	if err != nil {
		return 0, fmt.Errorf("Collection(%s).Documents().Search: %w", collectionName, err)
	}
	if result.Hits == nil {
		return 0, nil
	}

	list := make([]*map[string]any, 0, len(*result.Hits))
	for _, hit := range *result.Hits {
		list = append(list, hit.Document)
	}
	data, err := json.Marshal(list)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal: %w", err)
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return *result.OutOf, nil
}
