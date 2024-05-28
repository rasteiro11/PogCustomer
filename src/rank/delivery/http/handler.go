package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogPaymentSheet/src/rank"
	"net/http"
)

var RankGroupPath = "/rank"

type (
	HandlerOpt func(*handler)
	handler    struct {
		repository rank.Repository
	}
)

func WithRepository(repository rank.Repository) HandlerOpt {
	return func(u *handler) {
		u.repository = repository
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/list", RankGroupPath, http.MethodPost, h.List)
}

var ErrNotAuthorized = errors.New("not authorized")

func (h *handler) List(c *fiber.Ctx) error {
	creds, err := h.repository.List(c.Context())
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(creds))
}
