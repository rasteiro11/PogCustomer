package repository

import (
	"context"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/frequency"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ frequency.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) frequency.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Find(ctx context.Context, department *models.Frequency) (*models.Frequency, error) {
	res := &models.Frequency{}

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Where(department).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[frequency.repository.Find] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) Store(ctx context.Context, frequencia *models.Frequency) (*models.Frequency, error) {
	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Create(frequencia).Error; err != nil {
		logger.Of(ctx).Errorf("[frequency.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return frequencia, nil
}

func (r *repository) List(ctx context.Context) ([]models.Frequency, error) {
	var frequencia []models.Frequency

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Find(&frequencia).Error; err != nil {
		logger.Of(ctx).Errorf("[frequency.repository.FindAll] db.Find() returned error: %+v\n", err)
		return nil, err
	}

	return frequencia, nil
}

func (r *repository) UpdateById(ctx context.Context, user *models.Frequency) (*models.Frequency, error) {
	if err := r.db.Debug().Where(&models.Frequency{
		ID: user.ID,
	}).Updates(user).Error; err != nil {
		logger.Of(ctx).Errorf("[frequency.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return user, nil
}
