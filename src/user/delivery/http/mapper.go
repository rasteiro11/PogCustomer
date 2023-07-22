package http

import "github.com/rasteiro11/PogCustomer/models"

func loginRequestMapper(req *loginRequest) *models.User {
	return &models.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func loginResponseMapper(req *models.LoginResponse) *loginResponse {
	return &loginResponse{
		Token:     req.Token,
		ExpiresAt: req.ExpiresAt,
	}
}

func registerRequestMapper(req *registerRequest) *models.User {
	return &models.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func registerResponseMapper(req *models.RegisterResponse) *registerResponse {
	return &registerResponse{
		Token:     req.Token,
		ExpiresAt: req.ExpiresAt,
	}
}
