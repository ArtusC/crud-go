package service

import (
	"github.com/ArtusC/crud-go/src/configuration/logger"
	restErr "github.com/ArtusC/crud-go/src/configuration/restErr"
	"github.com/ArtusC/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(clientId string, userDomain model.UserDomainInterface) *restErr.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	return nil
}
