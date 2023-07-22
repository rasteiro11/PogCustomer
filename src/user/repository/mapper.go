package repository

import "github.com/rasteiro11/PogCustomer/models"

func userEntityMapper(query *models.User) *User {
	return &User{
		Email:    query.Email,
		Password: query.Password,
	}
}

func userMapper(query *User) *models.User {
	return &models.User{
		ID:       query.ID,
		Email:    query.Email,
		Password: query.Password,
	}
}
