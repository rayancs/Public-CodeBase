package cfg

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type JwtMetaData struct {
	Issuer     string
	Audience   string
	Secret     string
	Expiration int64
}

/*this is to load up the configuration throughout the application*/
func JwtConfig() *JwtMetaData {
	err := godotenv.Load()
	if err != nil {
		panic("could not load jwt configuration")
	}
	return &JwtMetaData{
		Issuer:     os.Getenv("JWT_ISSUER"),
		Audience:   os.Getenv("JWT_AUDIENCE"),
		Secret:     os.Getenv("JWT_SECRET"),
		Expiration: time.Now().Add(24 * time.Hour).Unix(),
	}
}
