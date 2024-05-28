package repository

import (
	"context"

	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/rank"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ rank.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) rank.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Find(ctx context.Context, rank *models.Rank) (*models.Rank, error) {
	res := &models.Rank{}

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Where(rank).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[rank.repository.Find] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) List(ctx context.Context) ([]models.Rank, error) {
	var ranks []models.Rank

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Find(&ranks).Error; err != nil {
		logger.Of(ctx).Errorf("[rank.repository.FindAll] db.Find() returned error: %+v\n", err)
		return nil, err
	}

	return ranks, nil
}
