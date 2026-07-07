package services

import (
	"context"
	"morng-dev/internal/core/domain/entities"
)

type AuthService interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error)
	Login(Ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error)
}
