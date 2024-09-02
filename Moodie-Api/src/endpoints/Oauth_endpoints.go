package endpoints

import (
	"moody/moody-api/src/handlers"

	"github.com/gin-gonic/gin"
)

/*this is the oauth end poit */
func OauthEndpoints(router *gin.Engine) {
	// get function
	//params explained :
	//	*gin.Engine.func (...params). where params = 1-address 2-handlers // side note :the order of functions are followed , exapmple (route , jwtMiddleware ,handler1 , handler2 ) it will be folllowed in order
	router.POST(GoogleEndPoint(), handlers.GoogleCallBackHandler)
}
func GoogleEndPoint() string { return "auth/google/callback" }
