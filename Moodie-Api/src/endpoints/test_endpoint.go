package endpoints

import (
	"moody/moody-api/src/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)
// this is only to test middlewares etc 
func TestEndPoint(r *gin.Engine) {
	r.GET("/test", middlewares.JwtMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "This End Point is PRotected By JWT")
	})
}
