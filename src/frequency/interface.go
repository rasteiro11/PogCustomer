package frequency

import (
	"context"
	"github.com/rasteiro11/PogPaymentSheet/models"
)

type (
	Repository interface {
		Find(ctx context.Context, department *models.Frequency) (*models.Frequency, error)
		Store(ctx context.Context, department *models.Frequency) (*models.Frequency, error)
		List(ctx context.Context) ([]models.Frequency, error)
		UpdateById(ctx context.Context, user *models.Frequency) (*models.Frequency, error)
	}
)
