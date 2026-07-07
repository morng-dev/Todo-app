package repositories

import (
	"context"
	"morng-dev/internal/core/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User, password string) error
	Getall(ctx context.Context, page, limit int) ([]*entities.User, int, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetPasswordHash(ctx context.Context, id uuid.UUID) (string, error)
}
