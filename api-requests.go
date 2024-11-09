package gotodoistsdk

import "net/http"

// struct with params for oauth request
type AuthorizationRequest struct {
	client_id string
	scope     string
	state     string
}

// struct with params for oauth response
type AuthorizationResponse struct {
	code  string
	state string
}

// struct with params for token exchange
type TokenRequest struct {
	client_id     string
	client_secret string
	code          string
}

// struct with access token and its type
type TokenResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

type Client struct {
	client        *http.Client
	client_id     string
	client_secret string
}

func NewClient(client *http.Client, clientId, clientSecret string) *Client {
	return &Client{client, clientId, clientSecret}
}
