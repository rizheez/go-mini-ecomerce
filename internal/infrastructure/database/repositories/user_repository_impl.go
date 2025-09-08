package repositories

import (
	"context"
	"mini-ecommerce/internal/domain/entities"
	"mini-ecommerce/internal/domain/repositories"
	"mini-ecommerce/internal/infrastructure/database/models"
	"time"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) GetById(ctx context.Context, id int) (*entities.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return toEntity(&user), nil
}

func (r *userRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return toEntity(&user), nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entities.User) error {
	userModels := &models.User{
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		Role:      user.Role,
		Phone:     user.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := r.db.WithContext(ctx).Create(userModels).Error; err != nil {
		return err
	}
	user.ID = userModels.ID
	return nil

}

func (r *userRepositoryImpl) Update(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(models.User{}, id).Error
}

func toEntity(user *models.User) *entities.User {
	return &entities.User{
		ID:            user.ID,
		Email:         user.Email,
		Name:          user.Name,
		Password:      user.Password,
		Role:          user.Role,
		Phone:         user.Phone,
		EmailVerified: user.EmailVerified,
		IsActive:      user.IsActive,
	}
}
