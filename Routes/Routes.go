package Routes

import (
	"github.com/gin-gonic/gin"
	"main/Controller"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/user/:id", Controller.GetUser)
	router.GET("/users/", Controller.GetUsers)
	router.POST("/user/:id", Controller.CreateUser)
	router.PUT("/user/:id", Controller.UpdateUser)
	router.DELETE("/user/:id", Controller.DeleteUser)
}
