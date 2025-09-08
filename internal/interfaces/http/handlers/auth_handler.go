package handlers

import (
	"mini-ecommerce/internal/interfaces/http/dto"
	"mini-ecommerce/internal/usecases"
	"mini-ecommerce/pkg/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	authUseCase usecases.AuthUsecase
}

// Login implements AuthHandler.
func (a *authHandler) Login(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Register implements AuthHandler.
func (a *authHandler) Register(c *fiber.Ctx) error {
	var req dto.UserReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
		})
	}
	if err := validation.Validate.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		errMsg := make(map[string]string)
		for _, e := range errs {
			field := e.Field()
			errMsg[field] = e.Translate(validation.Trans)
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": "Unprocessable Entity",
			"errors":  errMsg,
		})
	}
	res, err := a.authUseCase.Register(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error() + "asds",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data": &dto.UserRes{
			ID:        res.ID,
			Name:      res.Name,
			Email:     res.Email,
			Phone:     res.Phone,
			Role:      res.Role,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	})
}

func NewAuthHandler(authUseCase usecases.AuthUsecase) AuthHandler {
	return &authHandler{
		authUseCase: authUseCase,
	}
}
