package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/enums/usertype"
	"shared/pb/userpb"
)

type UpdateUserType struct {
	userClient userpb.UserServiceClient
}

func NewUpdateUserType(userClient userpb.UserServiceClient) *UpdateUserType {
	return &UpdateUserType{
		userClient: userClient,
	}
}

type UpdateUserTypeRequest struct {
	Id       string `params:"id"`
	UserType string `json:"user_type" validate:"required,oneof=admin viewer"`
}

func (h *UpdateUserType) Handle(ctx context.Context, req *UpdateUserTypeRequest) (*any, int, error) {
	userType := usertype.FromString(req.UserType)
	if !userType.IsValid() {
		return nil, http.StatusBadRequest, errors.New("invalid user type")
	}
	_, err := h.userClient.UpdateUserType(ctx, &userpb.UpdateUserTypeRequest{
		Id:       req.Id,
		UserType: userType.String(),
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("userClient.UpdateUserType: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
