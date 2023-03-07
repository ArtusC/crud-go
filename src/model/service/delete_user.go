package service

import (
	"github.com/ArtusC/crud-go/src/configuration/logger"
	rest_err "github.com/ArtusC/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
