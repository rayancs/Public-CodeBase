package service

import (
	"moody/moody-api/src/cfg"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignJwt(info map[string]interface{}) (string, error) {
	// ...
	jwtConfiguration := cfg.JwtConfig()
	signingKey := []byte(jwtConfiguration.Secret)

	subject := info["id"].(string) // Assuming 'id' is a field in the info map
	email := info["email"].(string)
	expirationTime := time.Now().Add(24 * time.Hour).Unix() // Token expires in 24 hours
	issuedAt := time.Now().Unix()

	// Create the JWT claims
	claims := jwt.MapClaims{
		"iss":   jwtConfiguration.Issuer,   // Issuer
		"sub":   subject,                   // Subject (typically the user ID)
		"aud":   jwtConfiguration.Audience, // Audience
		"exp":   expirationTime,            // Expiration time (in Unix seconds)
		"iat":   issuedAt,                  // Issued at (in Unix seconds)
		"email": email,                     // User's email address
		// Add any additional claims as needed
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil

}
