package repositories

import (
	"context"
	"morng-dev/internal/core/domain/entities"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Create(ctx context.Context, userID uuid.UUID, todo *entities.Todo) (*entities.Todo, error)
	GetById(ctx context.Context, id uuid.UUID) (*entities.Todo, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Todo, int, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	Delete(ctx context.Context, id uuid.UUID) error
}
