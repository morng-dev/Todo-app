package services

import (
	"context"
	"morng-dev/internal/core/domain/entities"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error)
	Login(Ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error)
	Logout(ctx context.Context, userID uuid.UUID) error
	RefresToken(ctx context.Context, req *entities.RefreshTokenRequest) (*entities.LoginResponse, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, req *entities.ChangePassword) error
	ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error
}
