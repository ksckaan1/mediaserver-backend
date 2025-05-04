package couchbasedb

import (
	"context"
	"errors"
	"fmt"
	"time"
	"user_service/internal/core/app"
	"user_service/internal/core/customerrors"
	"user_service/internal/core/models"

	"github.com/couchbase/gocb/v2"
	"github.com/samber/lo"
)

var _ app.Repository = (*Repository)(nil)

type Repository struct {
	coll  *gocb.Collection
	scope *gocb.Scope
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		coll:  bucket.Scope("user_service").Collection("users"),
		scope: bucket.Scope("user_service"),
	}
}

func (r *Repository) CreateUser(ctx context.Context, user *models.User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := r.coll.Insert(user.ID, user, &gocb.InsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("collection.Insert: %w", err)
	}
	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID string) error {
	_, err := r.coll.Remove(userID, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrUserNotFound
		}
		return fmt.Errorf("collection.Remove: %w", err)
	}
	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	result, err := r.coll.Get(id, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("collection.Get: %w", err)
	}
	var user models.User
	err = result.Content(&user)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &user, nil
}

const getUserByUsernameQuery = `SELECT * FROM users WHERE username = $username;`

type getUserByUsernameResult struct {
	Users models.User `json:"users"`
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	result, err := r.scope.Query(getUserByUsernameQuery, &gocb.QueryOptions{
		NamedParameters: map[string]any{
			"username": username,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	var user getUserByUsernameResult
	err = result.One(&user)
	if err != nil {
		if errors.Is(err, gocb.ErrNoResult) {
			return nil, customerrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("result.One: %w", err)
	}
	return &user.Users, nil
}

const (
	countUsersQuery         = `SELECT COUNT(*) AS count FROM users;`
	listUsersQuery          = `SELECT * FROM users OFFSET $offset;`
	listUsersWithLimitQuery = `SELECT * FROM users LIMIT $limit OFFSET $offset;`
)

type countResult struct {
	Count int64 `json:"count"`
}

type listResult struct {
	Users models.User `json:"users"`
}

func (r *Repository) ListUsers(ctx context.Context, limit int64, offset int64) (*models.UserList, error) {
	countCursor, err := r.scope.Query(countUsersQuery, &gocb.QueryOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	var countResult countResult
	err = countCursor.One(&countResult)
	if err != nil {
		return nil, fmt.Errorf("countCursor.One: %w", err)
	}
	if countResult.Count == 0 || limit == 0 {
		return &models.UserList{
			List:   []*models.User{},
			Count:  countResult.Count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	cursor, err := r.scope.Query(lo.Ternary(limit < 0, listUsersQuery, listUsersWithLimitQuery), &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"limit":  limit,
			"offset": offset,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	var users []*models.User
	for cursor.Next() {
		var user listResult
		err = cursor.Row(&user)
		if err != nil {
			return nil, fmt.Errorf("cursor.Row: %w", err)
		}
		users = append(users, &user.Users)
	}
	return &models.UserList{
		List:   users,
		Count:  countResult.Count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

const updatePasswordQuery = "UPDATE users SET `password` = $password WHERE id = $id RETURNING *;"

func (r *Repository) UpdateUserPassword(ctx context.Context, user *models.User) error {
	result, err := r.scope.Query(updatePasswordQuery, &gocb.QueryOptions{
		NamedParameters: map[string]any{
			"password": user.Password,
			"id":       user.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrUserNotFound
	}
	return nil
}

const updateUserTypeQuery = "UPDATE users SET `user_type` = $user_type WHERE id = $id RETURNING *;"

func (r *Repository) UpdateUserType(ctx context.Context, user *models.User) error {
	result, err := r.scope.Query(updateUserTypeQuery, &gocb.QueryOptions{
		NamedParameters: map[string]any{
			"user_type": user.UserType.String(),
			"id":        user.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrUserNotFound
	}
	return nil
}
