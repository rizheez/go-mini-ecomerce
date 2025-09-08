package usecases

import (
	"context"
	"errors"
	"mini-ecommerce/internal/domain/entities"
	"mini-ecommerce/internal/domain/repositories"
	"mini-ecommerce/internal/infrastructure/auth"
	"mini-ecommerce/internal/interfaces/http/dto"
	"time"
)

type AuthUsecase interface {
	Login(ctx context.Context, req *dto.UserLoginReq) (*dto.UserLoginRes, error)
	Register(ctx context.Context, req *dto.UserReq) (*dto.UserRes, error)
}
type authUseCaseImpl struct {
	userRepo repositories.UserRepository
}

// Login implements AuthUsecase.
func (a *authUseCaseImpl) Login(ctx context.Context, req *dto.UserLoginReq) (*dto.UserLoginRes, error) {
	panic("unimplemented")
}

// Register implements AuthUsecase.
func (a *authUseCaseImpl) Register(ctx context.Context, req *dto.UserReq) (*dto.UserRes, error) {
	exists, err := a.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if err.Error() != "user not found" {
			return nil, err
		}
	} else if exists != nil {
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
	err = a.userRepo.Create(ctx, user)
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

func NewAuthUsecase(userRepo repositories.UserRepository) AuthUsecase {
	return &authUseCaseImpl{
		userRepo: userRepo,
	}
}
