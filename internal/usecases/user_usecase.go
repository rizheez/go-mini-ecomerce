package usecases

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/entities"
	"mini-ecommerce/internal/domain/repositories"
	"mini-ecommerce/internal/infrastructure/auth"
	"mini-ecommerce/internal/interfaces/http/dto"
	"mini-ecommerce/pkg/logger"
	"time"
)

type UserUsecase interface {
	Login(ctx context.Context, req *dto.UserLoginReq) (*dto.UserLoginRes, error)
	GetById(ctx context.Context, id int) (*dto.UserRes, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserRes, error)
	Create(ctx context.Context, req *dto.UserReq) (*dto.UserRes, error)
	Update(ctx context.Context, req *dto.UserReq) (*dto.UserRes, error)
	Delete(ctx context.Context, id int) error
}

type userUseCaseImpl struct {
	userRepo repositories.UserRepository
}

// Create implements UserUsecase.
func (u *userUseCaseImpl) Create(ctx context.Context, req *dto.UserReq) (*dto.UserRes, error) {
	exists, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		logger.Error(errors.New("unprocessable Entity"), "[ErrUserUsecase-1] email already exists")
		return nil, errors.New("email already exists")
	}
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &entities.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: hashedPassword,
		Phone:    req.Phone,
		Role:     req.Role,
	}
	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &dto.UserRes{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil

}

// Delete implements UserUsecase.
func (u *userUseCaseImpl) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetByEmail implements UserUsecase.
func (u *userUseCaseImpl) GetByEmail(ctx context.Context, email string) (*dto.UserRes, error) {
	panic("unimplemented")
}

// GetById implements UserUsecase.
func (u *userUseCaseImpl) GetById(ctx context.Context, id int) (*dto.UserRes, error) {
	panic("unimplemented")
}

// Login implements UserUsecase.
func (u *userUseCaseImpl) Login(ctx context.Context, req *dto.UserLoginReq) (*dto.UserLoginRes, error) {
	panic("unimplemented")
}

// Update implements UserUsecase.
func (u *userUseCaseImpl) Update(ctx context.Context, req *dto.UserReq) (*dto.UserRes, error) {
	panic("unimplemented")
}

func NewUserUseCase(userRepo repositories.UserRepository) UserUsecase {
	return &userUseCaseImpl{
		userRepo: userRepo,
	}
}
