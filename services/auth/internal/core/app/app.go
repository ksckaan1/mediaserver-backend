package app

import (
	"auth_service/internal/core/models"
	"context"
	"fmt"
	"shared/pb/authpb"
	"shared/pb/userpb"
	"shared/ports"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ authpb.AuthServiceServer = (*App)(nil)

type App struct {
	authpb.UnimplementedAuthServiceServer
	userClient  userpb.UserServiceClient
	repository  Repository
	password    ports.Password
	idGenerator ports.IDGenerator
}

func New(userClient userpb.UserServiceClient, repository Repository, password ports.Password, idGenerator ports.IDGenerator) *App {
	return &App{
		userClient:  userClient,
		password:    password,
		idGenerator: idGenerator,
		repository:  repository,
	}
}

func (a *App) GetSession(ctx context.Context, req *authpb.GetSessionRequest) (*authpb.GetSessionResponse, error) {
	session, err := a.repository.GetSession(ctx, req.SessionId)
	if err != nil {
		return nil, fmt.Errorf("repository.GetSession: %w", err)
	}
	return &authpb.GetSessionResponse{
		SessionId: session.SessionID,
		CreatedAt: timestamppb.New(session.CreatedAt),
		UserId:    session.UserID,
		UserAgent: session.UserAgent,
		IpAddress: session.IpAddress,
	}, nil
}

func (a *App) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	user, err := a.userClient.GetUserByUsername(ctx, &userpb.GetUserByUsernameRequest{Username: req.Username})
	if err != nil {
		return nil, fmt.Errorf("userClient.GetUserByUsername: %w", err)
	}

	isPasswordValid, err := a.password.VerifyPassword(req.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("password.VerifyPassword: %w", err)
	}
	if !isPasswordValid {
		return nil, fmt.Errorf("password does not match")
	}

	sessionId := a.idGenerator.NewID()

	session := &models.Session{
		SessionID: sessionId,
		UserID:    user.Id,
		UserAgent: req.UserAgent,
		IpAddress: req.IpAddress,
	}

	err = a.repository.CreateSession(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateSession: %w", err)
	}
	return &authpb.LoginResponse{
		SessionId: session.SessionID,
	}, nil
}

func (a *App) Logout(ctx context.Context, req *authpb.LogoutRequest) (*emptypb.Empty, error) {
	err := a.repository.DeleteSession(ctx, req.SessionId)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteSession: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (a *App) LogoutAll(ctx context.Context, req *authpb.LogoutAllRequest) (*emptypb.Empty, error) {
	err := a.repository.DeleteAllSessionsByUserID(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteAllSessionsByUserID: %w", err)
	}
	return &emptypb.Empty{}, nil
}
