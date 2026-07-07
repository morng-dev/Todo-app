package services

import (
	"context"
	"fmt"
	"math"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/repositories"
	"morng-dev/internal/core/domain/ports/services"

	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo repositories.TodoRepository
}

func NewTodoRepository(todoRepo repositories.TodoRepository) services.TodoService {
	return &TodoService{todoRepo: todoRepo}
}

func (s *TodoService) CreateTodo(ctx context.Context, userID uuid.UUID, req *entities.TodoRequest) (*entities.Todo, error) {
	todo := &entities.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
	todo, err := s.todoRepo.Create(ctx, userID, todo)
	if err != nil {
		return nil, err
	}
	return todo, err
}

func (s *TodoService) GetAllTodo(ctx context.Context, page, limit int) ([]*entities.Todo, *entities.PaginationResponse, error) {
	todos, total, err := s.todoRepo.GetAll(ctx, page, limit)
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
	return todos, pagination, nil
}

func (s *TodoService) GetTodoByID(ctx context.Context, id uuid.UUID) (*entities.Todo, error) {
	return s.todoRepo.GetById(ctx, id)
}

func (s *TodoService) UpdateStatusTodo(ctx context.Context, userID, id uuid.UUID, req *entities.TodoStatusRequest) error {
	todo, err := s.GetTodoByID(ctx, id)
	if err != nil {
		return err
	}
	if todo.UserID != userID {
		return fmt.Errorf("คุณไม่สามารถแก้ไขได้")
	}
	return s.todoRepo.UpdateStatus(ctx, id, req.Status)
}

func (s *TodoService) DeleteTodo(ctx context.Context, userID, id uuid.UUID) error {
	todo, err := s.GetTodoByID(ctx, id)
	if err != nil {
		return err
	}
	if todo.UserID != userID {
		return fmt.Errorf("คุณไม่สามารถลบได้")
	}
	return s.todoRepo.Delete(ctx, id)
}
