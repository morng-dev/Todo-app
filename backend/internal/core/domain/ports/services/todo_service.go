package services

import (
	"context"
	"morng-dev/internal/core/domain/entities"

	"github.com/google/uuid"
)

type TodoService interface {
	CreateTodo(ctx context.Context, userID uuid.UUID, req *entities.TodoRequest) (*entities.Todo, error)
	GetTodoByID(ctx context.Context, id uuid.UUID) (*entities.Todo, error)
	GetAllTodo(ctx context.Context, page, limit int) ([]*entities.Todo, *entities.PaginationResponse, error)
	UpdateStatusTodo(ctx context.Context, userID, id uuid.UUID, status *entities.TodoStatusRequest) error
	DeleteTodo(ctx context.Context, userID, id uuid.UUID) error
}
