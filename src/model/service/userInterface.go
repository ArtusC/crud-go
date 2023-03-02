package service

import (
	restErr "github.com/ArtusC/crud-go/src/configuration/restErr"
	"github.com/ArtusC/crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *restErr.RestErr
	UpdateUser(string, model.UserDomainInterface) *restErr.RestErr
	FindUser(string) (*model.UserDomainInterface, *restErr.RestErr)
	DeleteUser(string) *restErr.RestErr
}
