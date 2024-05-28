package defaultvr

import (
	"context"
	"github.com/rasteiro11/PogPaymentSheet/models"
)

type (
	Repository interface {
		First(ctx context.Context, role *models.DefaultVR) (*models.DefaultVR, error)
	}
)
