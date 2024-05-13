package Routes

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.GET("/user/:id")
	router.GET("/users/")
	router.POST("/user/:id")
	router.PUT("/user/:id")
	router.DELETE("/user/:id")
}
