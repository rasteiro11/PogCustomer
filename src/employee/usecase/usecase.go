package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/defaultvr"
	"github.com/rasteiro11/PogPaymentSheet/src/department"
	"github.com/rasteiro11/PogPaymentSheet/src/employee"
	"github.com/rasteiro11/PogPaymentSheet/src/employeefrequency"
	"github.com/rasteiro11/PogPaymentSheet/src/frequency"
	"github.com/rasteiro11/PogPaymentSheet/src/rank"
	"github.com/rasteiro11/PogPaymentSheet/src/vr"
)

type (
	UsecaseOpt func(*usecase)
	usecase    struct {
		repository            employee.Repository
		departmentRepo        department.Repository
		defaultVrRepo         defaultvr.Repository
		vrRepo                vr.Repository
		rankRepo              rank.Repository
		employeeFrequencyRepo employeefrequency.Repository
		frequencyRepo         frequency.Repository
	}
)

var (
	ErrEmailTaken = errors.New("email already taken")
)

var _ employee.Usecase = (*usecase)(nil)

func WithEmployeeFrequencyRepo(repository employeefrequency.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.employeeFrequencyRepo = repository
	}
}

func WithFrequencyRepo(repository frequency.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.frequencyRepo = repository
	}
}

func WithRepository(repository employee.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.repository = repository
	}
}

func WithDefaultVrRepository(repository defaultvr.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.defaultVrRepo = repository
	}
}

func WithVrRepository(repository vr.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.vrRepo = repository
	}
}

func WithRankRepository(repository rank.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.rankRepo = repository
	}
}

func WithDepartmentRepository(repository department.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.departmentRepo = repository
	}
}

func NewUsecase(opts ...UsecaseOpt) employee.Usecase {
	u := &usecase{}

	for _, opt := range opts {
		opt(u)
	}

	return u
}

func (u *usecase) Promote(ctx context.Context, req *models.PromoteEmployeeRequest) (*models.Rank, error) {
	employee, err := u.repository.FindOne(ctx, &models.Employee{
		Name: req.EmployeeName,
	})
	if err != nil {
		return nil, err
	}

	newRank, err := u.rankRepo.Find(ctx, &models.Rank{
		Name: req.RankName,
	})
	if err != nil {
		return nil, err
	}

	employee.RankID = newRank.ID
	if _, err := u.repository.UpdateById(ctx, employee); err != nil {
		return nil, err
	}

	return newRank, nil
}

func getFirstDateOfCurrentMonth() time.Time {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	location := now.Location()

	firstDate := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, location)
	return firstDate
}

func (u *usecase) Add(ctx context.Context, req *models.CreateEmployeeRequest) (*models.Employee, error) {
	department, err := u.departmentRepo.Find(ctx, &models.Department{
		Name: req.DepartmentName,
	})
	if err != nil {
		return nil, err
	}

	defaultVr, err := u.defaultVrRepo.First(ctx, &models.DefaultVR{
		DepartmentID: department.ID,
	})
	if err != nil {
		return nil, err
	}

	createdVr, err := u.vrRepo.Store(ctx, &models.VR{
		Value: defaultVr.Value,
	})
	if err != nil {
		return nil, err
	}

	rank, err := u.rankRepo.Find(ctx, &models.Rank{
		Name: req.RankName,
	})
	if err != nil {
		return nil, err
	}

	employee, err := u.repository.Create(ctx, &models.Employee{
		Name:         req.Name,
		DepartmentID: department.ID,
		VrID:         createdVr.ID,
		Salary:       req.Salary,
		RankID:       rank.ID,
	})
	if err != nil {
		return nil, err
	}

	start := getFirstDateOfCurrentMonth()
	employeeFrequency, err := u.frequencyRepo.Store(ctx, &models.Frequency{
		DataReferencia: start,
		TotalDias:      0,
	})
	if err != nil {
		return nil, err
	}

	if _, err := u.employeeFrequencyRepo.Store(ctx, &models.EmployeeFrequency{
		DataReferencia: start,
		FuncionarioID:  employee.ID,
		FrequencyID:    employeeFrequency.ID,
	}); err != nil {
		return nil, err
	}

	return employee, nil
}

func (u *usecase) Delete(ctx context.Context, req *models.DeleteEmployeeRequest) (*models.Employee, error) {
	employee, err := u.repository.FindOne(ctx, &models.Employee{
		Name: req.EmployeeName,
	})
	if err != nil {
		return nil, err
	}

	if err := u.repository.Delete(ctx, &models.Employee{
		ID: employee.ID,
	}); err != nil {
		return nil, err
	}

	return employee, nil
}

func (u *usecase) List(ctx context.Context) ([]models.Employee, error) {
	return u.repository.FindAll(ctx)
}

func (u *usecase) UpdateById(ctx context.Context, user *models.UpdateEmployeeRequest) (*models.Employee, error) {
	employee, err := u.repository.FindOne(ctx, &models.Employee{
		Name: user.Name,
	})
	if err != nil {
		return nil, err
	}

	employee.Salary = user.Value
	return u.repository.UpdateById(ctx, employee)
}
