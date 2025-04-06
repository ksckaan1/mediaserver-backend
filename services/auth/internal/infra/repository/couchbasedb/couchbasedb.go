package couchbasedb

import (
	"auth_service/internal/core/customerrors"
	"auth_service/internal/core/models"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

type Repository struct {
	coll  *gocb.Collection
	scope *gocb.Scope
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		coll:  bucket.Scope("auth_service").Collection("sessions"),
		scope: bucket.Scope("auth_service"),
	}
}

func (r *Repository) CreateSession(ctx context.Context, session *models.Session) error {
	session.CreatedAt = time.Now()
	_, err := r.coll.Upsert(session.SessionID, session, &gocb.UpsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("coll.Upsert: %w", err)
	}
	return nil
}

func (r *Repository) GetSession(ctx context.Context, sessionId string) (*models.Session, error) {
	result, err := r.coll.Get(sessionId, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrSessionNotFound
		}
		return nil, fmt.Errorf("coll.Get: %w", err)
	}
	session := &models.Session{}
	err = result.Content(session)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return session, nil
}

func (r *Repository) DeleteSession(ctx context.Context, sessionId string) error {
	_, err := r.coll.Remove(sessionId, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrSessionNotFound
		}
		return fmt.Errorf("coll.Remove: %w", err)
	}
	return nil
}

const deleteAllSessionsByUserIDQuery = `DELETE FROM sessions WHERE user_id = $user_id;`

func (r *Repository) DeleteAllSessionsByUserID(ctx context.Context, userId string) error {
	_, err := r.scope.Query(deleteAllSessionsByUserIDQuery, &gocb.QueryOptions{
		Context:         ctx,
		NamedParameters: map[string]any{"user_id": userId},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	return nil
}
