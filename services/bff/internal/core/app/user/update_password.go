package user

import (
	"bff-service/internal/pkg/sessionutils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/enums/usertype"
	"shared/pb/userpb"
	"shared/ports"
)

type UpdatePassword struct {
	userClient userpb.UserServiceClient
	password   ports.Password
}

func NewUpdatePassword(userClient userpb.UserServiceClient, password ports.Password) *UpdatePassword {
	return &UpdatePassword{
		userClient: userClient,
		password:   password,
	}
}

type UpdatePasswordRequest struct {
	Id                 string `params:"id"`
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password" validate:"required,min=8,max=32"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8,max=32,eqfield=NewPassword"`
}

func (h *UpdatePassword) Handle(ctx context.Context, req *UpdatePasswordRequest) (*any, int, error) {
	session, err := sessionutils.GetSession(ctx)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("session not found")
	}
	if !(session.UserType == usertype.Admin || session.UserId == req.Id) {
		return nil, http.StatusUnauthorized, errors.New("only admin or owner can update password")
	}

	if session.UserType != usertype.Admin {
		user, err := h.userClient.GetUserByID(ctx, &userpb.GetUserByIDRequest{
			Id: req.Id,
		})
		if err != nil {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		ok, err := h.password.VerifyPassword(req.OldPassword, user.Password)
		if err != nil {
			return nil, http.StatusUnauthorized, errors.New("invalid password")
		}
		if !ok {
			return nil, http.StatusUnauthorized, errors.New("old passwords do not match")
		}
	}
	if req.NewPassword != req.ConfirmNewPassword {
		return nil, http.StatusBadRequest, errors.New("new passwords do not match")
	}
	hashedPassword, err := h.password.HashPassword(req.NewPassword)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("password.HashPassword: %w", err)
	}
	_, err = h.userClient.UpdateUserPassword(ctx, &userpb.UpdateUserPasswordRequest{
		Id:       req.Id,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("userClient.UpdateUserPassword: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
