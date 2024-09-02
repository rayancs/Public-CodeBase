// this will get the email from gin context , that is extracted from the jwt token
package service

import "github.com/gin-gonic/gin"

func GetEmailFromContext(ctx *gin.Context) string {
	email, exists := ctx.Get("email")
	if !exists {
		return "" // Return an empty string if the email does not exist in the context
	}
	return email.(string) // Return the email as a string
}
