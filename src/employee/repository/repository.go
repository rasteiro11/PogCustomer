package repository

import (
	"context"

	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/employee"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ employee.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) employee.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Delete(ctx context.Context, employee *models.Employee) error {
	if err := r.db.Delete(employee).Where(employee).Error; err != nil {
		logger.Of(ctx).Errorf("[employee.repository.Delete] db.Take() returned error: %+v\n", err)
		return err
	}

	return nil
}

func (r *repository) FindOne(ctx context.Context, user *models.Employee) (*models.Employee, error) {
	res := &models.Employee{}
	if err := r.db.Where(user).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[employee.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) Create(ctx context.Context, user *models.Employee) (*models.Employee, error) {
	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Create(user).Error; err != nil {
		logger.Of(ctx).Errorf("[employee.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return user, nil
}

func (r *repository) FindAll(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Find(&employees).Error; err != nil {
		logger.Of(ctx).Errorf("[employees.repository.FindAll] db.Find() returned error: %+v\n", err)
		return nil, err
	}

	return employees, nil
}

func (r *repository) UpdateById(ctx context.Context, user *models.Employee) (*models.Employee, error) {
	if err := r.db.Debug().Where(&models.Employee{
		ID: user.ID,
	}).Updates(user).Error; err != nil {
		logger.Of(ctx).Errorf("[employee.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return user, nil
}

func (r *repository) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		ctx := database.WithTx(ctx, tx)
		return fn(ctx)
	})
}
