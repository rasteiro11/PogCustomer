package repository

import (
	"context"

	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/vr"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ vr.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) vr.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Find(ctx context.Context, vr *models.VR) (*models.VR, error) {
	res := &models.VR{}

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Where(vr).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[vr.repository.Find] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) Store(ctx context.Context, vr *models.VR) (*models.VR, error) {
	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Create(vr).Error; err != nil {
		logger.Of(ctx).Errorf("[vr.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return vr, nil
}

func (r *repository) UpdateById(ctx context.Context, user *models.VR) (*models.VR, error) {
	if err := r.db.Debug().Where(&models.VR{
		ID: user.ID,
	}).Updates(user).Error; err != nil {
		logger.Of(ctx).Errorf("[vr.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return user, nil
}
