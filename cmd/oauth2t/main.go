package main

import (
	"context"
	"fmt"

	"github.com/Contrast-Security-Inc/oauth2t/internal/providers"
	"github.com/coreos/go-oidc"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var oauthConfig clientcredentials.Config
var provider oidc.Provider

func processConfig() {
	viper.SetDefault("provider", "auth0")
	viper.SetDefault("issuer", "https://cs-tenant-1.us.auth0.com/")
	viper.SetDefault("audience", "http://harmony-api")

	viper.SetEnvPrefix("oauth2t")
	viper.BindEnv("client_id")
	viper.BindEnv("client_secret")
	viper.BindEnv("provider")
	viper.BindEnv("issuer")
	viper.BindEnv("audenice")

	viper.SetConfigName("oauth2t")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("config file not found\n")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func NewTokenSource(issuer, clientID, clientSecret string) oauth2.TokenSource {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		panic(err)
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauthConfig = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		// Discovery returns the OAuth2 endpoints.
		TokenURL: provider.Endpoint().TokenURL,
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return oauthConfig.TokenSource(context.Background())
}

func getTokenSource() oauth2.TokenSource {

	issuer := viper.GetString("issuer")
	clientID := viper.GetString("client_id")
	clientSecret := viper.GetString("client_secret")
	provider := viper.GetString("provider")
	audience := viper.GetString("audience")

	switch provider {
	case "auth0":
		return providers.NewAuth0TokenSource(issuer, clientID, clientSecret, audience)
	default:
		return NewTokenSource(issuer, clientID, clientSecret)
	}
}

func main() {
	processConfig()
	ts := getTokenSource()
	token, err := ts.Token()
	if err != nil {
		panic(err)
	}
	fmt.Printf(token.AccessToken)
}
