package handlers

import (
	"moody/moody-api/src/models"
	"moody/moody-api/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this is the handler for the user api
func GetAllUserHandler(context *gin.Context) {
	// get the user id from the request

	context.JSON(http.StatusOK, "t-1")

}
func PostSingleUser(context *gin.Context) {
	// Create the user
	// json validation using gin
	var user models.UserRequestModel
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad Json Data"})
		return
	}
	//validate operation
	isSucc, strRes := service.CreateNewUser(user)
	if !isSucc {
		context.JSON(http.StatusBadRequest, gin.H{"error": strRes})
		return
	}
	context.JSON(http.StatusOK, strRes)

}
