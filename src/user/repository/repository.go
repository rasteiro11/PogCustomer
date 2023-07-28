package repository

import (
	"context"
	"github.com/rasteiro11/PogCore/pkg/logger"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ user.Repository = (*repository)(nil)

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewRepository(db *gorm.DB) user.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) FindOne(ctx context.Context, user *models.User) (*models.User, error) {
	res := &User{}
	if err := r.db.Where(userEntityMapper(user)).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[user.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return userMapper(res), nil
}

func (r *repository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	res := userEntityMapper(user)
	if err := r.db.Create(res).Error; err != nil {
		logger.Of(ctx).Errorf("[user.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return user, nil
}

func (r *repository) FindOneByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	res := &User{}
	if err := r.db.Where(&User{Email: user.Email}).Take(res).Error; err != nil {
		logger.Of(ctx).Errorf("[user.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return userMapper(res), nil
}

func (r *repository) UpdateById(ctx context.Context, user *models.User) (*models.User, error) {
	query := userEntityMapper(user)
	if err := r.db.Debug().Where(User{
		Model: gorm.Model{
			ID: user.ID,
		},
	}).Updates(query).Error; err != nil {
		logger.Of(ctx).Errorf("[user.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return userMapper(query), nil
}
