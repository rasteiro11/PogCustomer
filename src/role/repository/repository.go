package repository

import (
	"context"

	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/role"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ role.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) role.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Find(ctx context.Context, role *models.Role) (*models.Role, error) {
	res := &models.Role{}

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Where(role).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[role.repository.Find] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) Store(ctx context.Context, role *models.Role) (*models.Role, error) {
	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Create(role).Error; err != nil {
		logger.Of(ctx).Errorf("[role.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return role, nil
}
