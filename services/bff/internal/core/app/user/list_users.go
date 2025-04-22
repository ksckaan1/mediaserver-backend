package user

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/enums/usertype"
	"shared/pb/userpb"

	"github.com/samber/lo"
)

type ListUsers struct {
	userClient userpb.UserServiceClient
}

func NewListUsers(userClient userpb.UserServiceClient) *ListUsers {
	return &ListUsers{userClient: userClient}
}

type ListUsersRequest struct {
	Limit  *int64 `query:"limit"`
	Offset int64  `query:"offset"`
}

type ListUsersResponse struct {
	List   []*models.User `json:"list"`
	Count  int64          `json:"count"`
	Limit  int64          `json:"limit"`
	Offset int64          `json:"offset"`
}

func (h *ListUsers) Handle(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, int, error) {
	limit := int64(10)
	if req.Limit != nil {
		limit = *req.Limit
	}
	users, err := h.userClient.ListUsers(ctx, &userpb.ListUsersRequest{
		Limit:  limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("userClient.ListUsers: %w", err)
	}
	return &ListUsersResponse{
		List: lo.Map(users.List, func(user *userpb.User, _ int) *models.User {
			user.Password = ""
			return &models.User{
				ID:        user.Id,
				CreatedAt: user.CreatedAt.AsTime(),
				UpdatedAt: user.UpdatedAt.AsTime(),
				Username:  user.Username,
				Password:  user.Password,
				UserType:  usertype.FromString(user.UserType),
			}
		}),
		Count:  users.Count,
		Limit:  limit,
		Offset: req.Offset,
	}, http.StatusOK, nil
}
