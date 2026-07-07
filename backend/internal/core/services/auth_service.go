package services

import (
	"context"
	"errors"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/repositories"
	"morng-dev/internal/core/domain/ports/services"
	"morng-dev/pkg/utils"
)

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) services.AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error) {
	if _, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return nil, errors.New("อีเมลนี้มีอยู่ในระบบแล้ว")
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &entities.User{
		Email: req.Email,
		Name:  req.Name,
	}
	if err := s.userRepo.Create(ctx, user, hashPassword); err != nil {
		return nil, err
	}
	return s.userRepo.GetByID(ctx, user.ID)
}

func (s *authService) Login(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("อีเมลหรือรหัสผ่านไม่ถูกต้อง")
	}

	hashPassword, err := s.userRepo.GetPasswordHash(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(hashPassword, req.Password) {
		return nil, errors.New("อีเมลหรือรหัสผ่านไม่ถูกต้อง")
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Email)
	if err != nil {
		return nil, err
	}
	return &entities.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
