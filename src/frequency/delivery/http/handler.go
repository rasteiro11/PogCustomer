package http

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/employee"
	"github.com/rasteiro11/PogPaymentSheet/src/employeefrequency"
	"github.com/rasteiro11/PogPaymentSheet/src/frequency"
)

var FrequencyGroupPath = "/frequency"

type (
	HandlerOpt func(*handler)
	handler    struct {
		repository            frequency.Repository
		employeeFrequencyRepo employeefrequency.Repository
		employeeRepo          employee.Repository
	}
)

func WithRepository(repository frequency.Repository) HandlerOpt {
	return func(u *handler) {
		u.repository = repository
	}
}

func WithEmployeeFrequencyRepo(repository employeefrequency.Repository) HandlerOpt {
	return func(u *handler) {
		u.employeeFrequencyRepo = repository
	}
}

func WithEmployeeRepo(repository employee.Repository) HandlerOpt {
	return func(u *handler) {
		u.employeeRepo = repository
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/get", FrequencyGroupPath, http.MethodPost, h.Get)
	server.AddHandler("/update", FrequencyGroupPath, http.MethodPost, h.Update)
}

var ErrNotAuthorized = errors.New("not authorized")

func (h *handler) Update(c *fiber.Ctx) error {
	req := &models.UpdateFrequencyRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	employee, err := h.employeeRepo.FindOne(c.Context(), &models.Employee{
		Name: req.EmployeeName,
	})
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	employeeFrequency, err := h.employeeFrequencyRepo.Find(c.Context(), &models.EmployeeFrequency{
		DataReferencia: req.Date,
		FuncionarioID:  employee.ID,
	})
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	creds, err := h.repository.Find(c.Context(), &models.Frequency{
		ID:             employeeFrequency.FrequencyID,
		DataReferencia: req.Date,
	})
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	creds.TotalDias = req.Frequency
	updateFrequency, err := h.repository.UpdateById(c.Context(), creds)
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(updateFrequency))
}

func (h *handler) Get(c *fiber.Ctx) error {
	req := &models.GetFrequencyRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	employee, err := h.employeeRepo.FindOne(c.Context(), &models.Employee{
		Name: req.EmployeeName,
	})
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	employeeFrequency, err := h.employeeFrequencyRepo.Find(c.Context(), &models.EmployeeFrequency{
		DataReferencia: req.Date,
		FuncionarioID:  employee.ID,
	})
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	creds, err := h.repository.Find(c.Context(), &models.Frequency{
		ID:             employeeFrequency.FrequencyID,
		DataReferencia: req.Date,
	})
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(creds))
}
