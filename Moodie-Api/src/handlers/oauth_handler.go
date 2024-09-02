package handlers

import (
	"context"
	"moody/moody-api/src/cfg"
	"moody/moody-api/src/literals"
	"moody/moody-api/src/models"
	"moody/moody-api/src/service"
	"moody/moody-api/src/service/google_oauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this is the main login function aswell , so this is the requierments for this function
/*1-obtain a the access token to get user information like email
2-check if user exsisted or not if not then create it in your own server , email , id (google.UUID)
3-create a jwt by server , conatining the email

*/

// this is the handler for oauth feature
func GoogleCallBackHandler(c *gin.Context) {
	// the code from the client is used to make a http GET request the googles end point (auth flow ,where you use a code to get access token and refresh token  )
	code := c.Query("code")
	// using the code and all of the extra details like client id , client secret etc to generate a request token
	//!!! explination!!!//
	// This is done by making an HTTP POST request to the token endpoint of the OAuth provider,
	//  providing details such as the authorization code, client ID, client secret, and the redirect URI.
	token, err := cfg.OAuthConfig().Exchange(context.Background(), code)
	if err != nil {
		// replace error with error response struct
		noTokenCtrErr := models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Error:   []string{literals.GeneralError("Logging You In")},
			Details: []string{"no_token_gen", "auth_flow_err", err.Error()},
		}
		c.JSON(noTokenCtrErr.Code, noTokenCtrErr)
		return
	}
	// this is from the service , more explination over there
	data, err := google_oauth.GetUserInfo(token.AccessToken)
	if err != nil {
		userInfoErr := models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Error:   []string{literals.GeneralError("Logging You In")},
			Details: []string{"no_token_gen", "auth_flow_err", err.Error()},
		}
		c.JSON(userInfoErr.Code, userInfoErr)
		return
	}
	// create the user over here
	userPresistance := models.UserRequestModel{
		Email: data["email"].(string),
	}
	isSuccess, err := service.CreateNewUser(userPresistance)
	if err != nil {
		userCreationError := models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Error:   []string{literals.GeneralError("Logging You In")},
			Details: []string{"user_creation_fail", err.Error()},
		}
		c.JSON(userCreationError.Code, userCreationError)
		return
	}
	if !isSuccess {
		useCreationFail := models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Error:   []string{literals.GeneralError("Logging You In")},
			Details: []string{"no_success_while_login_process"},
		}
		c.JSON(useCreationFail.Code, useCreationFail)
		return
	}

	// server generated JWT token
	// this is from the service , more explination there
	responseToken, err := service.SignJwt(data)
	if err != nil {
		jwtTokenGenErr := models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Error:   []string{literals.GeneralError("Logging You In")},
			Details: []string{"jwt_token_gen_err", "service_err", err.Error()},
		}
		c.JSON(jwtTokenGenErr.Code, jwtTokenGenErr)
		return
	}
	// create a better response
	c.JSON(200, responseToken)

}
