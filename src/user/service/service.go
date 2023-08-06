package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rasteiro11/PogCore/pkg/config"
	pbCustomer "github.com/rasteiro11/PogCustomer/gen/proto/go/customer"
	"github.com/rasteiro11/PogCustomer/models"
	"github.com/rasteiro11/PogCustomer/src/user"
)

type service struct {
	userUsecase user.Usecase
}

var (
	ErrInvalidToken     = errors.New("error invalid token")
	ErrSignatureInvalid = errors.New("error invalid signature")
)

type Option func(*service)

var _ pbCustomer.CustomerServiceServer = (*service)(nil)

func WithUserUsecase(userUsecase user.Usecase) Option {
	return func(s *service) {
		s.userUsecase = userUsecase
	}
}

func NewService(opts ...Option) pbCustomer.CustomerServiceServer {
	s := &service{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *service) GetUser(ctx context.Context, req *pbCustomer.GetUserRequest) (*pbCustomer.GetUserResponse, error) {
	user, err := s.userUsecase.FindOne(ctx, userMapper(req))
	if err != nil {
		return nil, err
	}

	return getUserResponseMapper(user), nil
}

func (s *service) VerifySession(ctx context.Context, req *pbCustomer.VerifySessionRequest) (*pbCustomer.VerifySessionResponse, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(req.Token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Instance().RequiredString("JWT_SECRET")), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, ErrSignatureInvalid
		}
		return nil, err
	}

	if !token.Valid {
		return nil, ErrSignatureInvalid
	}

	return &pbCustomer.VerifySessionResponse{
		UserId: uint64(claims.UserID),
	}, nil
}
