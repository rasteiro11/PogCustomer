package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogCustomer/src/user"
	"net/http"
	"time"
)

var AuthGroupPath = "/auth"

type (
	HandlerOpt func(*handler)
	handler    struct {
		usecase user.Usecase
	}
)

func WithUsecase(usecase user.Usecase) HandlerOpt {
	return func(u *handler) {
		u.usecase = usecase
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/signin", AuthGroupPath, http.MethodPost, h.Login)
	server.AddHandler("/register", AuthGroupPath, http.MethodPost, h.Register)
	server.AddHandler("/changepassword", AuthGroupPath, http.MethodPost, h.ChangePassword)
}

var ErrNotAuthorized = errors.New("not authorized")

var _ user.Handler = (*handler)(nil)

type loginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type loginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type registerRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Document string `json:"document"`
}

type registerResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type changePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
	Email       string `json:"email"`
}

type changePasswordResponse struct {
}

func (h *handler) Login(c *fiber.Ctx) error {
	req := &loginRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	creds, err := h.usecase.Login(c.Context(), loginRequestMapper(req))
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(loginResponseMapper(creds)))
}

func (h *handler) Register(c *fiber.Ctx) error {
	req := &registerRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	creds, err := h.usecase.Register(c.Context(), registerRequestMapper(req))
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusCreated(c, rest.WithBody(creds))
}

func (h *handler) ChangePassword(c *fiber.Ctx) error {
	req := &changePasswordRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	creds, err := h.usecase.ChangePassword(c.Context(), changePasswordRequestMapper(req))
	if err != nil {
		return rest.NewStatusUnprocessableEntity(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(changePasswordResponseMapper(creds)))
}
