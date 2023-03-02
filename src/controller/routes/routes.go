package routes

import (
	controller "github.com/ArtusC/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {

	// recebe o parâmetro usrerIds
	r.GET("/getUserById/:userId", userController.FindUserByID)

	// recebe o parâmetro usrerEmail
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)

	// não recebe parâmetro
	r.POST("/createUser", userController.CreateUser)

	// recebe o parâmetro usrerId
	r.PUT("/updateUser/:userId", userController.UpdateUser)

	// recebe o parâmetro usrerId
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
