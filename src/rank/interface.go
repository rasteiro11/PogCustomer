package rank

import (
	"context"
	"github.com/rasteiro11/PogPaymentSheet/models"
)

type (
	Repository interface {
		Find(ctx context.Context, rank *models.Rank) (*models.Rank, error)
		List(ctx context.Context) ([]models.Rank, error)
	}
)
