package user

import (
	"bff-service/internal/pkg/sessionutils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/enums/usertype"
	"shared/pb/userpb"
	"time"
)

type Profile struct {
	userClient userpb.UserServiceClient
}

func NewProfile(userClient userpb.UserServiceClient) *Profile {
	return &Profile{
		userClient: userClient,
	}
}

type ProfileRequest struct{}

type ProfileResponse struct {
	ID        string            `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Username  string            `json:"username"`
	UserType  usertype.UserType `json:"user_type"`
}

func (h *Profile) Handle(ctx context.Context, req *ProfileRequest) (*ProfileResponse, int, error) {
	session, err := sessionutils.GetSession(ctx)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("session not found")
	}

	user, err := h.userClient.GetUserByID(ctx, &userpb.GetUserByIDRequest{
		Id: session.UserId,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("userClient.GetUserByID: %w", err)
	}

	return &ProfileResponse{
		ID:        user.Id,
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
		Username:  user.Username,
		UserType:  usertype.FromString(user.UserType),
	}, http.StatusOK, nil
}
