package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogCustomer/models"
)

type (
	Handler interface {
		Login(c *fiber.Ctx) error
		Register(c *fiber.Ctx) error
	}
	Usecase interface {
		Login(ctx context.Context, req *models.User) (*models.LoginResponse, error)
		Register(ctx context.Context, req *models.User) (*models.RegisterResponse, error)
	}
	Repository interface {
		FindOne(ctx context.Context, user *models.User) (*models.User, error)
		Create(ctx context.Context, user *models.User) (*models.User, error)
		FindOneByEmail(ctx context.Context, user *models.User) (*models.User, error)
	}
)
