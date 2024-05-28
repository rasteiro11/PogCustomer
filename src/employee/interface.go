package employee

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogPaymentSheet/models"
)

type (
	Handler interface {
		Add(c *fiber.Ctx) error
	}
	Usecase interface {
		Add(ctx context.Context, req *models.CreateEmployeeRequest) (*models.Employee, error)
		Delete(ctx context.Context, req *models.DeleteEmployeeRequest) (*models.Employee, error)
		List(ctx context.Context) ([]models.Employee, error)
		Promote(ctx context.Context, req *models.PromoteEmployeeRequest) (*models.Rank, error)
		UpdateById(ctx context.Context, user *models.UpdateEmployeeRequest) (*models.Employee, error)
	}
	Repository interface {
		FindOne(ctx context.Context, user *models.Employee) (*models.Employee, error)
		Create(ctx context.Context, user *models.Employee) (*models.Employee, error)
		Delete(ctx context.Context, employee *models.Employee) error
		FindAll(ctx context.Context) ([]models.Employee, error)
		UpdateById(ctx context.Context, user *models.Employee) (*models.Employee, error)
		Tx(ctx context.Context, fn func(ctx context.Context) error) error
	}
)
