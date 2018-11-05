package easycron

import (
	"net/http"
)

// Config for the Easycron client.
type Config struct {
	HTTPClient  *http.Client
	AccessToken string
	BaseURL     string
}

// NewConfig with the given access token.
func NewConfig(accessToken string) *Config {
	return &Config{
		HTTPClient:  http.DefaultClient,
		AccessToken: accessToken,
		BaseURL:     "https://www.easycron.com/rest/",
	}
}
