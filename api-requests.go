package gotodoistsdk

import "net/http"

// AuthorizationRequest struct with params for oauth request
type AuthorizationRequest struct {
	clientId string
	scope    string
	state    string
}

// AuthorizationResponse struct with params for oauth response
type AuthorizationResponse struct {
	code  string
	state string
}

// TokenRequest struct with params for token exchange
type TokenRequest struct {
	clientId     string
	clientSecret string
	code         string
}

// TokenResponse struct with access token and its type
type TokenResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

type Client struct {
	client       *http.Client
	clientId     string
	clientSecret string
}

func NewClient(client *http.Client, clientId, clientSecret string) *Client {
	return &Client{client, clientId, clientSecret}
}
