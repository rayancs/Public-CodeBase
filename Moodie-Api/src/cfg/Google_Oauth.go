package cfg

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

/*this is to load up the configuration throughout the application*/
func OAuthConfig() *oauth2.Config {
	err := godotenv.Load()
	if err != nil {
		panic("error loading oauth credentials")
	}
	return &oauth2.Config{
		ClientID: os.Getenv("G_CLIENT_ID"),
		ClientSecret: os.Getenv("G_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("G_REDIRECT_URL"),
		Scopes:       []string{"profile", "email"},
	}

}
