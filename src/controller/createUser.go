package controller

import (
	"net/http"

	"github.com/ArtusC/crud-go/src/configuration/logger"
	"github.com/ArtusC/crud-go/src/configuration/validation"
	"github.com/ArtusC/crud-go/src/configuration/view"
	"github.com/ArtusC/crud-go/src/controller/model/request"
	"github.com/ArtusC/crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marshal object", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// como o objeto model é privado, o único jeito de acessá-lo será
	// pelo uso do construtor abaixo. Dessa forma, garantimos que
	// nenhum outro método modifique as informações do usuário.
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age)

	// service := service.NewUserDomainService()
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))

}
