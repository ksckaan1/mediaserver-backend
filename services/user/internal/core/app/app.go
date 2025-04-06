package app

import (
	"context"
	"errors"
	"fmt"
	"shared/pb/userpb"
	"shared/ports"
	"user_service/internal/core/customerrors"
	"user_service/internal/core/models"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ userpb.UserServiceServer = (*App)(nil)

type App struct {
	repository Repository
	userpb.UnimplementedUserServiceServer
	idGenerator ports.IDGenerator
	password    ports.Password
}

func New(repository Repository, idGenerator ports.IDGenerator, password ports.Password) *App {
	return &App{
		repository:  repository,
		idGenerator: idGenerator,
		password:    password,
	}
}

func (a *App) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	_, err := a.repository.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return nil, customerrors.ErrUsernameAlreadyInUse
	} else if !errors.Is(err, customerrors.ErrUserNotFound) {
		return nil, fmt.Errorf("repository.GetUserByUsername: %w", err)
	}

	id := a.idGenerator.NewID()

	hashedPassword, err := a.password.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("password.HashPassword: %w", err)
	}

	user := &models.User{
		ID:       id,
		Username: req.Username,
		Password: hashedPassword,
	}

	err = a.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateUser: %w", err)
	}

	return &userpb.CreateUserResponse{
		UserId: id,
	}, nil
}

func (a *App) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	err := a.repository.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteUser: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (a *App) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.User, error) {
	user, err := a.repository.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("repository.GetUserByID: %w", err)
	}
	return &userpb.User{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		Password:  user.Password,
	}, nil
}

func (a *App) GetUserByUsername(ctx context.Context, req *userpb.GetUserByUsernameRequest) (*userpb.User, error) {
	user, err := a.repository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("repository.GetUserByUsername: %w", err)
	}
	return &userpb.User{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		Password:  user.Password,
	}, nil
}

func (a *App) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userList, err := a.repository.ListUsers(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("repository.ListUsers: %w", err)
	}
	return &userpb.ListUsersResponse{
		List: lo.Map(userList.List, func(user *models.User, _ int) *userpb.User {
			return &userpb.User{
				Id:        user.ID,
				Username:  user.Username,
				CreatedAt: timestamppb.New(user.CreatedAt),
				UpdatedAt: timestamppb.New(user.UpdatedAt),
				Password:  user.Password,
			}
		}),
		Count:  userList.Count,
		Limit:  userList.Limit,
		Offset: userList.Offset,
	}, nil
}

func (a *App) UpdateUserPassword(ctx context.Context, req *userpb.UpdateUserPasswordRequest) (*emptypb.Empty, error) {
	hashedPassword, err := a.password.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("password.HashPassword: %w", err)
	}
	err = a.repository.UpdateUserPassword(ctx, &models.User{
		ID:       req.Id,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, fmt.Errorf("repository.UpdateUserPassword: %w", err)
	}
	return &emptypb.Empty{}, nil
}
