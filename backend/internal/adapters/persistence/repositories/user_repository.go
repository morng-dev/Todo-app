package repositories

import (
	"context"
	"morng-dev/internal/adapters/persistence/models"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User, password string) error {
	userModel := &models.User{
		Email:    user.Email,
		Password: password,
		Name:     user.Name,
	}

	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return err
	}
	user.ID = userModel.ID
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var userModel models.User
	if err := r.db.WithContext(ctx).First(&userModel, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return r.modelsToEntities(&userModel), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var userModel models.User

	if err := r.db.WithContext(ctx).First(&userModel, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return r.modelsToEntities(&userModel), nil
}
func (r *UserRepository) Getall(ctx context.Context, page, limit int) ([]*entities.User, int, error) {
	var users []models.User
	var total int64
	offest := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Preload("Todo").Offset(offest).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	var result []*entities.User

	for _, user := range users {
		result = append(result, r.modelsToEntities(&user))
	}
	return result, int(total), nil
}

func (r *UserRepository) GetPasswordHash(ctx context.Context, id uuid.UUID) (string, error) {
	var user models.User

	if err := r.db.WithContext(ctx).Select("password").First(&user, "id = ?", id).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}

func (r *UserRepository) modelsToEntities(userModel *models.User) *entities.User {
	user := &entities.User{
		ID:        userModel.ID,
		Email:     userModel.Email,
		Name:      userModel.Name,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}

	return user
}
