package service

import (
	pbCustomer "github.com/rasteiro11/PogCustomer/gen/proto/go/customer"
	"github.com/rasteiro11/PogCustomer/models"
)

func userMapper(req *pbCustomer.GetUserRequest) *models.User {
	return &models.User{
		ID: uint(req.Id),
	}
}

func getUserResponseMapper(user *models.User) *pbCustomer.GetUserResponse {
	return &pbCustomer.GetUserResponse{
		Id:       int32(user.ID),
		Email:    user.Email,
		Document: user.Document,
	}
}

func getUserByDocumentRequest(req *pbCustomer.GetUserByDocumentRequest) *models.User {
	return &models.User{
		Document: req.Document,
	}
}

func getUserByDocumentResponse(req *models.User) *pbCustomer.GetUserByDocumentResponse {
	return &pbCustomer.GetUserByDocumentResponse{
		Id:       int32(req.ID),
		Email:    req.Email,
		Document: req.Document,
	}
}
