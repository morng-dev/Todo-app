package repositories

import (
	"context"
	"morng-dev/internal/adapters/persistence/models"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repositories.TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, userID uuid.UUID, todo *entities.Todo) (*entities.Todo, error) {
	todoModel := models.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Status:      "pending",
		UserID:      userID,
	}
	if err := r.db.WithContext(ctx).Create(&todoModel).Error; err != nil {
		return nil, err
	}
	return r.modelsToEntities(&todoModel), nil
}

func (r *TodoRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.Todo, error) {
	var todomodel models.Todo
	if err := r.db.WithContext(ctx).Preload("User").First(&todomodel, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return r.modelsToEntities(&todomodel), nil
}
func (r *TodoRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Todo, int, error) {
	var todos []models.Todo
	var total int64

	offset := (page - 1) * limit
	if err := r.db.WithContext(ctx).Model(&models.Todo{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.WithContext(ctx).Preload("User").Offset(offset).Limit(limit).Find(&todos).Error; err != nil {
		return nil, 0, err
	}
	var result []*entities.Todo

	for _, todo := range todos {
		result = append(result, r.modelsToEntities(&todo))
	}

	return result, int(total), nil
}

func (r *TodoRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return r.db.WithContext(ctx).Model(&models.Todo{}).Where("id = ?", id).Update("status", status).Error
}

func (r *TodoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Todo{}, "id = ?", id).Error
}
func (r *TodoRepository) modelsToEntities(todomodel *models.Todo) *entities.Todo {
	var user *entities.User
	if todomodel.User != nil {
		user = &entities.User{
			ID:    todomodel.UserID,
			Email: todomodel.User.Email,
			Name:  todomodel.User.Name,
		}
	}
	return &entities.Todo{
		ID:          todomodel.ID,
		Title:       todomodel.Title,
		Description: todomodel.Description,
		Status:      todomodel.Status,
		UserID:      todomodel.UserID,
		User:        user,
		CreatedAt:   todomodel.CreatedAt,
		UpdatedAt:   todomodel.UpdatedAt,
	}
}
