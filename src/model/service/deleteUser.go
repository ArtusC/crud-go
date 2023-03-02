package service

import (
	"github.com/ArtusC/crud-go/src/configuration/logger"
	restErr "github.com/ArtusC/crud-go/src/configuration/restErr"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(string) *restErr.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
