package Controller

import (
	"github.com/gin-gonic/gin"
	"main/Configuration/rest_err"
	"net/http"
)

func CreateUser(context *gin.Context) {

}
func UpdateUser(context *gin.Context) {

}
func GetUsers(context *gin.Context) {

	err := rest_err.NewBadRequestError("bad request")
	context.JSON(http.StatusBadRequest, err)

}
func GetUser(context *gin.Context) {

}
func DeleteUser(context *gin.Context) {

}
