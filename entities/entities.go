package entities

import "github.com/rasteiro11/PogCustomer/models"

func GetEntities() []any {
	return []any{
		&models.User{},
		&models.Role{},
		&models.UserRole{},
		&models.DefaultRole{},
	}
}
