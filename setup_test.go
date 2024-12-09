package shiji_test

import (
	"context"
	"fmt"
	"os"

	"github.com/omniboost/go-shiji"
)

func client() *shiji.Client {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")
	shijiTenant := os.Getenv("SHIJI_TENANT")
	shijiUsername := os.Getenv("SHIJI_USERNAME")
	shijiPassword := os.Getenv("SHIJI_PASSWORD")

	oauthConfig := shiji.NewOauth2PasswordConfig()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret
	oauthConfig.Username = fmt.Sprintf("%s@%s", shijiUsername, shijiTenant)
	oauthConfig.Password = shijiPassword

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client := shiji.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)
	return client
}
