package entities

import "github.com/rasteiro11/PogPaymentSheet/models"

func GetEntities() []any {
	return []any{
		&models.VR{},
		&models.DefaultVR{},
		&models.Department{},
		&models.Rank{},
		&models.Employee{},
		&models.EmployeeFrequency{},
		&models.Frequency{},
	}
}
