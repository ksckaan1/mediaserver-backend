package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/userpb"
	"strings"
)

type Register struct {
	userClient userpb.UserServiceClient
}

func NewRegister(userClient userpb.UserServiceClient) *Register {
	return &Register{userClient: userClient}
}

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,alphanum,min=3,max=15"`
	Password        string `json:"password" validate:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=32,eqfield=Password"`
}

type RegisterResponse struct {
	UserId string `json:"user_id"`
}

func (r *Register) Handle(ctx context.Context, req *RegisterRequest) (*RegisterResponse, int, error) {
	resp, err := r.userClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		if strings.Contains(err.Error(), "username already in use") {
			return nil, http.StatusConflict, errors.New("username already in use")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("userClient.CreateUser: %w", err)
	}
	return &RegisterResponse{
		UserId: resp.UserId,
	}, http.StatusCreated, nil
}
