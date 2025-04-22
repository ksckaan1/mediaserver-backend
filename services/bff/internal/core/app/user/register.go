package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/enums/usertype"
	"shared/pb/userpb"
	"strings"

	"github.com/samber/lo"
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
	users, err := r.userClient.ListUsers(ctx, &userpb.ListUsersRequest{
		Limit:  0,
		Offset: 0,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("userClient.ListUsers: %w", err)
	}
	userType := lo.Ternary(users.Count == 0, usertype.Admin, usertype.Viewer)
	resp, err := r.userClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		UserType: userType.String(),
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
