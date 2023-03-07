package service

import (
	"github.com/ArtusC/crud-go/src/configuration/logger"
	rest_err "github.com/ArtusC/crud-go/src/configuration/rest_err"
	"github.com/ArtusC/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	return nil, nil
}
