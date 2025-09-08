package routes

import (
	"mini-ecommerce/internal/infrastructure/database/repositories"
	"mini-ecommerce/internal/interfaces/http/handlers"
	"mini-ecommerce/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := repositories.NewUserRepositoryImpl(db)
	authUseCase := usecases.NewAuthUsecase(userRepo)
	authHandler := handlers.NewAuthHandler(authUseCase)
	SetupAuthRoutes(app, authHandler)
}
