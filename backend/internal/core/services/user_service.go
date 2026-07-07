package services

import (
	"context"
	"math"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/repositories"
	"morng-dev/internal/core/domain/ports/services"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) services.UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUsers(ctx context.Context, page, limit int) ([]*entities.User, *entities.PaginationResponse, error) {
	users, total, err := s.userRepo.Getall(ctx, page, limit)
	if err != nil {
		return nil, nil, err
	}
	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	pagination := &entities.PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPage,
		TotalItems: total,
	}
	return users, pagination, nil
}
