package github.com/hhhanafubukiii/go-todoist-sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	client       *http.Client
	clientId     string
	clientSecret string
}

type Task struct {
	id          string
	name        string
	priority    string
	description string
	dueDate     string
}

type AddTask struct {
	name        string
	priority    string
	dueDate     string
	description string
}

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

func NewClient(client *http.Client, clientId, clientSecret string) *Client {
	return &Client{client, clientId, clientSecret}
}

func (t *Client) AddTask(
	name,
	priority,
	description,
	date,
	token string,
) error {
	response, err := t.createAddTaskRequest(name, priority, description, date, token)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	return nil
}

func (t *Client) createAddTaskRequest(
	name,
	priority,
	description,
	date,
	accessToken string,
) (*http.Response, error) {
	client := t.client

	requestURL := fmt.Sprintf(`%s/rest/%s/%s?content=%s&priority=%s&description=%s&due_string=%s`,
		BASE_URL,
		REST_VERSION,
		TASKS_ENDPOINT,
		name,
		priority,
		description,
		date,
	)

	requestBody := setAddTaskBody(name, priority, description, date)
	bodyReader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}

func setAddTaskBody(name, priority, description, date string) []byte {
	s := fmt.Sprintf(`
							{
							"content": %s
							"priority": %s
							"description": %s
							"date": %s
							}
							`, name, priority, description, date)
	body, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	return body
}