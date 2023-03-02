package controller

import (
	"github.com/ArtusC/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

// metodo para inciar o controller
func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

// instância de todos os services (para não precisar ficar criando toda vez que é chamado)
type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

// criando objeto para chamar a instância do service
type userControllerInterface struct {
	service service.UserDomainService
}
