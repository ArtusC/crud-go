package service

import (
	"github.com/ArtusC/crud-go/src/configuration/logger"
	rest_err "github.com/ArtusC/crud-go/src/configuration/rest_err"
	"github.com/ArtusC/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(clientId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	return nil
}
