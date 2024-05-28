package repository

import (
	"context"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/defaultrole"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ defaultrole.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) defaultrole.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) First(ctx context.Context, role *models.DefaultRole) (*models.DefaultRole, error) {
	res := &models.DefaultRole{}

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Where(role).First(res).Error; err != nil {
		logger.Of(ctx).Errorf("[defaultrole.repository.First] db.First() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}
