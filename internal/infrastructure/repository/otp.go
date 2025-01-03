package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/repository/sqlcgen"
)

func (r *Repository) CreateOTP(ctx context.Context, otp *model.OTP) error {
	err := r.queries.CreateOTP(ctx, sqlcgen.CreateOTPParams{
		Email:     otp.Email,
		Code:      otp.Code,
		ExpiresAt: otp.ExpiresAt,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateOTP: %w", err)
	}
	return nil
}

func (r *Repository) GetOTPByEmailAndCode(ctx context.Context, email, code string) (*model.OTP, error) {
	otp, err := r.queries.GetOTPByEmailAndCode(ctx, sqlcgen.GetOTPByEmailAndCodeParams{
		Email: email,
		Code:  code,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetOTPByEmailAndCode: %w", customerrors.ErrOTPNotFound)
		}
		return nil, fmt.Errorf("queries.GetOTPByEmailAndCode: %w", err)
	}
	return &model.OTP{
		Email:     otp.Email,
		Code:      otp.Code,
		CreatedAt: otp.CreatedAt,
		ExpiresAt: otp.ExpiresAt,
	}, nil
}

func (r *Repository) DeleteOTPByEmail(ctx context.Context, email string) error {
	err := r.queries.DeleteOTPByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("queries.DeleteOTPByEmail: %w", err)
	}
	return nil
}
