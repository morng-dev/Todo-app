package repositories

import (
	"context"
	"morng-dev/internal/adapters/persistence/models"
	"morng-dev/internal/core/domain/entities"
	"morng-dev/internal/core/domain/ports/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repositories.CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, userID, todoID uuid.UUID, req *entities.Comment) (*entities.Comment, error) {
	comment := models.Comment{
		Content: req.Content,
		TodoID:  todoID,
		UserID:  userID,
	}
	if err := r.db.WithContext(ctx).Create(&comment).Error; err != nil {
		return nil, err
	}
	return r.modelsToEntities(&comment), nil
}

func (r *CommentRepository) Update(ctx context.Context, commentID uuid.UUID, req *entities.Comment) error {
	updates := map[string]interface{}{}

	if req.Content != "" {
		updates["Content"] = req.Content
	}
	return r.db.WithContext(ctx).Model(&models.Comment{}).Where("id = ?", commentID).Updates(&updates).Error
}

func (r *CommentRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&models.Comment{}).Where("id = ?", commentID).Error
}

func (r *CommentRepository) modelsToEntities(commentModel *models.Comment) *entities.Comment {
	comments := &entities.Comment{
		ID:        commentModel.ID,
		Content:   commentModel.Content,
		TodoID:    commentModel.TodoID,
		UserID:    commentModel.UserID,
		CreatedAt: commentModel.CreatedAt,
		UpdatedAt: commentModel.UpdatedAt,
		DeletedAt: commentModel.DeletedAt,
	}
	if commentModel.UserID != uuid.Nil {
		comments.User = &entities.User{
			ID:    commentModel.UserID,
			Email: commentModel.User.Email,
			Name:  commentModel.User.Name,
		}
	}
	return comments
}
