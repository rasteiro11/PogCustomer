package http

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogPaymentSheet/src/department"
)

var DepartmentGroupPath = "/department"

type (
	HandlerOpt func(*handler)
	handler    struct {
		repository department.Repository
	}
)

func WithRepository(repository department.Repository) HandlerOpt {
	return func(u *handler) {
		u.repository = repository
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/list", DepartmentGroupPath, http.MethodPost, h.List)
}

var ErrNotAuthorized = errors.New("not authorized")

func (h *handler) List(c *fiber.Ctx) error {
	creds, err := h.repository.List(c.Context())
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(creds))
}
