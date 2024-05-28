package role

import (
	"context"
	"github.com/rasteiro11/PogCustomer/models"
)

type (
	Repository interface {
		Find(ctx context.Context, role *models.Role) (*models.Role, error)
		Store(ctx context.Context, role *models.Role) (*models.Role, error)
	}
)
