package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogCore/pkg/validator"
	"github.com/rasteiro11/PogPaymentSheet/models"
	"github.com/rasteiro11/PogPaymentSheet/src/employee"
	"net/http"
)

var EmployeeGroupPath = "/employee"

type (
	HandlerOpt func(*handler)
	handler    struct {
		usecase employee.Usecase
	}
)

func WithUsecase(usecase employee.Usecase) HandlerOpt {
	return func(u *handler) {
		u.usecase = usecase
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/add", EmployeeGroupPath, http.MethodPost, h.Add)
	server.AddHandler("/update", EmployeeGroupPath, http.MethodPost, h.Update)
	server.AddHandler("/delete", EmployeeGroupPath, http.MethodPost, h.Delete)
	server.AddHandler("/list", EmployeeGroupPath, http.MethodPost, h.List)
	server.AddHandler("/promote", EmployeeGroupPath, http.MethodPost, h.Promote)
}

var ErrNotAuthorized = errors.New("not authorized")

var _ employee.Handler = (*handler)(nil)

func (h *handler) Promote(c *fiber.Ctx) error {
	req := &models.PromoteEmployeeRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	if _, err := validator.IsRequestValid(req); err != nil {
		return rest.NewResponse(c, http.StatusBadRequest, rest.WithBody(err)).JSON(c)
	}

	creds, err := h.usecase.Promote(c.Context(), req)
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(creds))
}

func (h *handler) List(c *fiber.Ctx) error {
	creds, err := h.usecase.List(c.Context())
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(creds))
}

func (h *handler) Add(c *fiber.Ctx) error {
	req := &models.CreateEmployeeRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	if _, err := validator.IsRequestValid(req); err != nil {
		return rest.NewResponse(c, http.StatusBadRequest, rest.WithBody(err)).JSON(c)
	}

	creds, err := h.usecase.Add(c.Context(), req)
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(creds))
}

func (h *handler) Delete(c *fiber.Ctx) error {
	req := &models.DeleteEmployeeRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	if _, err := validator.IsRequestValid(req); err != nil {
		return rest.NewResponse(c, http.StatusBadRequest, rest.WithBody(err)).JSON(c)
	}

	deletedEmployee, err := h.usecase.Delete(c.Context(), req)
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(deletedEmployee))
}

func (h *handler) Update(c *fiber.Ctx) error {
	req := &models.UpdateEmployeeRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	employee, err := h.usecase.UpdateById(c.Context(), &models.UpdateEmployeeRequest{
		Name:  req.Name,
		Value: req.Value,
	})
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(employee))
}
