package http

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/employee"
	"github.com/rasteiro11/PogPaymentSheet/src/vr"
)

var VrGroupPath = "/vr"

type (
	HandlerOpt func(*handler)
	handler    struct {
		repository  vr.Repository
		emloyeeRepo employee.Repository
	}
)

func WithEmployeeRepo(repository employee.Repository) HandlerOpt {
	return func(u *handler) {
		u.emloyeeRepo = repository
	}
}

func WithRepository(repository vr.Repository) HandlerOpt {
	return func(u *handler) {
		u.repository = repository
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/find", VrGroupPath, http.MethodPost, h.Get)
	server.AddHandler("/update", VrGroupPath, http.MethodPost, h.Update)
}

var ErrNotAuthorized = errors.New("not authorized")

func (h *handler) Get(c *fiber.Ctx) error {
	req := &models.GetVrRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	creds, err := h.repository.Find(c.Context(), &models.VR{
		ID: req.VrID,
	})
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(creds))
}

func (h *handler) Update(c *fiber.Ctx) error {
	req := &models.UpdateVrRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	employee, err := h.emloyeeRepo.FindOne(c.Context(), &models.Employee{
		Name: req.EmployeeName,
	})
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	creds, err := h.repository.Find(c.Context(), &models.VR{
		ID: employee.VrID,
	})
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	creds.Value = req.Value
	updatedVr, err := h.repository.UpdateById(c.Context(), creds)
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(updatedVr))
}
