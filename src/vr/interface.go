package vr

import (
	"context"
	"github.com/rasteiro11/PogPaymentSheet/models"
)

type (
	Repository interface {
		Find(ctx context.Context, vr *models.VR) (*models.VR, error)
		Store(ctx context.Context, vr *models.VR) (*models.VR, error)
		UpdateById(ctx context.Context, user *models.VR) (*models.VR, error)
	}
)
