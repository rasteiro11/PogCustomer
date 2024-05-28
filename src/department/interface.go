package department

import (
	"context"
	"github.com/rasteiro11/PogPaymentSheet/models"
)

type (
	Repository interface {
		Find(ctx context.Context, department *models.Department) (*models.Department, error)
		Store(ctx context.Context, department *models.Department) (*models.Department, error)
		List(ctx context.Context) ([]models.Department, error)
	}
)
