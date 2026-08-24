package repositories

import (
	"context"
	"morng-dev/internal/core/domain/entities"

	"github.com/google/uuid"
)

type CommentRepository interface {
	Create(ctx context.Context, TodoID, UserID uuid.UUID, req *entities.Comment) (*entities.Comment, error)
	Update(ctx context.Context, commentID uuid.UUID, req *entities.Comment) error
	Delete(ctx context.Context, commentID uuid.UUID) error
}
