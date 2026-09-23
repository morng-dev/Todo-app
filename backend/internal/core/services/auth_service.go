package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/repositories"
	"morng-dev/internal/core/domain/ports/services"
	"morng-dev/pkg/utils"

	"github.com/google/uuid"
)

type authService struct {
	userRepo repositories.UserRepository
	mailer   services.Mailer
}

func NewAuthService(userRepo repositories.UserRepository, mailer services.Mailer) services.AuthService {
	return &authService{userRepo: userRepo, mailer: mailer}
}

func (s *authService) Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error) {
	exists, err := s.userRepo.GetEmailExist(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
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
		User:  user,
	}, nil
}

func (s *authService) Logout(ctx context.Context, userID uuid.UUID) error {
	return s.userRepo.SetRefreshToken(ctx, userID, "")
}

func (s *authService) ChangePassword(ctx context.Context, userID uuid.UUID, req *entities.ChangePassword) error {
	hashPassword, err := s.userRepo.GetPasswordHash(ctx, userID)
	if err != nil {
		return err
	}
	if !utils.CheckPassword(hashPassword, req.OldPassword) {
		return errors.New("password invalid")
	}
	newHashPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	return s.userRepo.UpdatePassword(ctx, userID, newHashPassword)
}

func (s *authService) ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest) error {
	exists, err := s.userRepo.GetEmailExist(ctx, req.Email)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	resetToken, err := s.generateResetToken()
	if err != nil {
		return err
	}
	if err := s.userRepo.SetResetToken(ctx, req.Email, resetToken); err != nil {
		return err
	}

	if err := s.mailer.SendResetPassword(ctx, req.Email, resetToken); err != nil {
		return errors.New("ส่งอีเมลไม่สำเร็จ กรุณาลองใหม่อีกครั้ง")
	}
	return nil
}

func (s *authService) RefresToken(ctx context.Context, req *entities.RefreshTokenRequest) (*entities.LoginResponse, error) {
	user, err := s.userRepo.GetByRefresh(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateJWT(user.ID.String(), user.Email)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}
	if err := s.userRepo.SetRefreshToken(ctx, user.ID, refreshToken); err != nil {
		return nil, err
	}
	return &entities.LoginResponse{
		Token:       token,
		RefresToken: refreshToken,
	}, nil
}

func (s *authService) ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error {
	user, err := s.userRepo.GetByResetToken(ctx, req.Token)
	if err != nil {
		return errors.New("invalid or expired reset token")
	}
	newHashPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	s.userRepo.UpdatePassword(ctx, user.ID, newHashPassword)
	return s.userRepo.ClearResetToken(ctx, user.ID)
}

func (s *authService) generateRefreshToken() (string, error) {
	byte := make([]byte, 32)
	if _, err := rand.Read(byte); err != nil {
		return "", err
	}
	return hex.EncodeToString(byte), nil
}

func (s *authService) generateResetToken() (string, error) {
	byte := make([]byte, 16)
	if _, err := rand.Read(byte); err != nil {
		return "", nil
	}
	return hex.EncodeToString(byte), nil

}
