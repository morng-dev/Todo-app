package services

import (
	"context"
	"morng-dev/internal/core/domain/entities"
)

type UserService interface {
	GetUsers(ctx context.Context, page, limit int) ([]*entities.User, *entities.PaginationResponse, error)
}
