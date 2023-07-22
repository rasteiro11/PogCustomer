package http

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rasteiro11/PogCore/pkg/server"
	"github.com/rasteiro11/PogCore/pkg/transport/rest"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/user"
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
	server.AddHandler("/welcome", AuthGroupPath, http.MethodPost, h.Welcome)
}

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
}

type registerResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
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
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusCreated(c, rest.WithBody(creds))
}

func (h *handler) Welcome(c *fiber.Ctx) error {
	jwtToken := c.GetReqHeaders()
	tok, ok := jwtToken["Authorization"]
	if !ok {
		return rest.NewStatusUnauthorized(c, errors.New("not authorized"))
	}

	tok = strings.ReplaceAll(tok, "Bearer ", "")
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tok, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return rest.NewStatusUnauthorized(c, err)
		}
		return rest.NewStatusBadRequest(c, err)
	}

	if !token.Valid {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(claims))
}
