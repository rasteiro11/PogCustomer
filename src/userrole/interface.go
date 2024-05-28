package userrole

import (
	"context"
	"github.com/rasteiro11/PogCustomer/models"
)

type (
	Repository interface {
		Find(ctx context.Context, role *models.UserRole) (*models.UserRole, error)
		Store(ctx context.Context, role *models.UserRole) (*models.UserRole, error)
	}
)
