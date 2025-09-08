package repositories

import (
	"context"
	"mini-ecommerce/internal/domain/entities"
)

type UserRepository interface {
	GetById(ctx context.Context, id int) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id int) error
}
