package service

import (
	"github.com/ArtusC/crud-go/src/configuration/logger"
	restErr "github.com/ArtusC/crud-go/src/configuration/restErr"
	"github.com/ArtusC/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *restErr.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	return nil
}
