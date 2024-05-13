package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Configuration/rest_err"
	"main/Model/UserModel"
)

func CreateUser(context *gin.Context) {
	var UserRequest UserModel.UserRequest
	err := context.ShouldBind(&UserRequest)
	if err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("there are some incorrect field, error= %s", err.Error()))
		context.JSON(restErr.Code, restErr)
	}
}

func UpdateUser(context *gin.Context) {

}
func GetUsers(context *gin.Context) {

}
func GetUser(context *gin.Context) {

}
func DeleteUser(context *gin.Context) {

}
