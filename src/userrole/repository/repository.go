package repository

import (
	"context"

	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/userrole"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ userrole.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) userrole.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Store(ctx context.Context, role *models.UserRole) (*models.UserRole, error) {
	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Create(role).Error; err != nil {
		logger.Of(ctx).Errorf("[userrole.repository.Create] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return role, nil
}

func (r *repository) Find(ctx context.Context, role *models.UserRole) (*models.UserRole, error) {
	res := &models.UserRole{}

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Where(role).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[userrole.repository.Find] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}
