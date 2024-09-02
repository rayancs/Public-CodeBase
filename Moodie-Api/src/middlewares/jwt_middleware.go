package middlewares

// detail comment all of the files uptil now
import (
	"fmt"
	"moody/moody-api/src/cfg"
	"moody/moody-api/src/literals"
	"moody/moody-api/src/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			errorRes := models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Error:   []string{"Authorization Could Not Be Succesful Due To Missing Credentials"},
				Details: []string{"null_token"},
			}
			c.JSON(errorRes.Code, errorRes)
			c.Abort()
			return
		}
		// the token is supposed to be in this form
		/*
			key , value
			Authorization , Bearer xyz
			the line below should provide you with just the token
		*/
		tokenExtract := strings.TrimPrefix(token, literals.Bearer())
		// this means the token is not provided properly
		if tokenExtract == token {
			errorRes := models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Error:   []string{"Authorization Could Not Be Succesful Due To InAppropriate Credentials"},
				Details: []string{"empty_auth_header"},
			}
			c.JSON(errorRes.Code, errorRes)
			c.Abort()
			return
		}
		// improt configurations
		jwtConfiguration := cfg.JwtConfig()

		secret := []byte(jwtConfiguration.Secret)
		// parse the token ?
		literals.RedLog(tokenExtract)
		parsedToken, err := jwt.Parse(tokenExtract, func(tkn *jwt.Token) (interface{}, error) {
			if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secret, nil
		})
		if err != nil {
			errorRes := models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Error:   []string{"Authorization Could Not Be Succesful Due To InAppropriate Credentials"},
				Details: []string{"token_not_parsed", err.Error()},
			}
			c.JSON(errorRes.Code, errorRes)
			c.Abort()
			return
		}
		if !parsedToken.Valid {
			errorRes := models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Error:   []string{"Authorization Could Not Be Succesful Due To InAppropriate Credentials"},
				Details: []string{"invalid_parsed_token"},
			}
			c.JSON(errorRes.Code, errorRes)
			c.Abort()
			return
		}
		// validate all of the claims
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			errorRes := models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Error:   []string{"Authorization Could Not Be Succesful Due To InAppropriate Credentials"},
				Details: []string{"claims_not_processed"},
			}
			c.JSON(http.StatusUnauthorized, errorRes)
			c.Abort()
			return
		}
		// verify if the claims match each other
		if claims["iss"] != jwtConfiguration.Issuer || claims["aud"] != jwtConfiguration.Audience {
			claimsError := models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Error:   []string{"Authorization Could Not Be Succesful Due To InAppropriate Credentials"},
				Details: []string{"claims_no_match"},
			}
			c.JSON(claimsError.Code, claimsError)
			c.Abort()
			// pass the email from the jwt token

			return
		}
		if email, ok := claims["email"].(string); ok {
			// this will set the email to the gin context
			c.Set("email", email)
		} else {
			errorRes := models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Error:   []string{"Authorization Could Not Be Succesful Due To InAppropriate Credentials"},
				Details: []string{"email_claim_not_found"},
			}
			c.JSON(http.StatusUnauthorized, errorRes)
			c.Abort()
			return
		}
		// pass the email in the context so that the next funtion can get the email of the user
		c.Next()
	}

}
