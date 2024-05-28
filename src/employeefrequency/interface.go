package employeefrequency

import (
	"context"
	"github.com/rasteiro11/PogPaymentSheet/models"
)

type (
	Repository interface {
		Find(ctx context.Context, department *models.EmployeeFrequency) (*models.EmployeeFrequency, error)
		Store(ctx context.Context, department *models.EmployeeFrequency) (*models.EmployeeFrequency, error)
	}
)
