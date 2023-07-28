package service

import (
	"context"

	pbCustomer "github.com/rasteiro11/PogCustomer/gen/proto/go/customer"
	"github.com/rasteiro11/PogCustomer/src/user"
)

type service struct {
	userUsecase user.Usecase
}

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
