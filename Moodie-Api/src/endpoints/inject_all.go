package endpoints

import "github.com/gin-gonic/gin"

/* this is for better code readability */
/*All of the end points are to be injected here */
func InjectEndpoints(eng *gin.Engine) {
	TestEndPoint(eng)
	OauthEndpoints(eng)

}
