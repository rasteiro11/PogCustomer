package repository

import (
	"context"
	"github.com/rasteiro11/PogCore/pkg/database"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/employeefrequency"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ employeefrequency.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) employeefrequency.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Find(ctx context.Context, department *models.EmployeeFrequency) (*models.EmployeeFrequency, error) {
	res := &models.EmployeeFrequency{}

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Where(department).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[vr.repository.Find] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) Store(ctx context.Context, department *models.EmployeeFrequency) (*models.EmployeeFrequency, error) {
	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Create(department).Error; err != nil {
		logger.Of(ctx).Errorf("[department.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return department, nil
}

func (r *repository) List(ctx context.Context) ([]models.EmployeeFrequency, error) {
	var departments []models.EmployeeFrequency

	tx, err := database.FromContext(ctx)
	if err != nil {
		tx = r.db
	}

	if err := tx.Debug().Find(&departments).Error; err != nil {
		logger.Of(ctx).Errorf("[departments.repository.FindAll] db.Find() returned error: %+v\n", err)
		return nil, err
	}

	return departments, nil
}
