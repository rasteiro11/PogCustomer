package defaultrole

import (
	"context"
	"github.com/rasteiro11/PogCustomer/models"
)

type (
	Repository interface {
		First(ctx context.Context, role *models.DefaultRole) (*models.DefaultRole, error)
	}
)
