package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/repository/sqlcgen"

	"github.com/samber/lo"
)

func (r *Repository) CreateUser(ctx context.Context, user *model.User) error {
	err := r.queries.CreateUser(ctx, sqlcgen.CreateUserParams{
		ID:          user.DisplayName,
		Email:       user.Email,
		DisplayName: user.DisplayName,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateUser: %w", err)
	}
	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetUserByID: %w", customerrors.ErrUserNotFound)
		}
		return nil, fmt.Errorf("queries.GetUserByID: %w", err)
	}
	return &model.User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		DisplayName: user.DisplayName,
	}, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetUserByEmail: %w", customerrors.ErrUserNotFound)
		}
		return nil, fmt.Errorf("queries.GetUserByEmail: %w", err)
	}
	return &model.User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		DisplayName: user.DisplayName,
	}, nil
}

func (r *Repository) ListUsers(ctx context.Context, limit, offset int64) (*model.UserList, error) {
	count, err := r.queries.CountUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.CountUsers: %w", err)
	}
	if count == 0 {
		return &model.UserList{
			Users:  make([]model.User, 0),
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	users, err := r.queries.ListUsers(ctx, sqlcgen.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("queries.ListUsers: %w", err)
	}
	return &model.UserList{
		Users: lo.Map(users, func(u sqlcgen.User, _ int) model.User {
			return model.User{
				ID:          u.ID,
				CreatedAt:   u.CreatedAt,
				UpdatedAt:   u.UpdatedAt,
				Email:       u.Email,
				DisplayName: u.DisplayName,
			}
		}),
	}, nil
}

func (r *Repository) UpdateUserByID(ctx context.Context, user *model.User) error {
	_, err := r.queries.UpdateUserByID(ctx, sqlcgen.UpdateUserByIDParams{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		ID:          user.ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.UpdateUser: %w", customerrors.ErrUserNotFound)
		}
		return fmt.Errorf("queries.UpdateUser: %w", err)
	}
	return nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id string) error {
	_, err := r.queries.DeleteUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.DeleteUserByID: %w", customerrors.ErrUserNotFound)
		}
		return fmt.Errorf("queries.DeleteUserByID: %w", err)
	}
	return nil
}
