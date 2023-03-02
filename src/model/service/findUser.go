package service

import (
	"github.com/ArtusC/crud-go/src/configuration/logger"
	restErr "github.com/ArtusC/crud-go/src/configuration/restErr"
	"github.com/ArtusC/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *restErr.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	return nil, nil
}
