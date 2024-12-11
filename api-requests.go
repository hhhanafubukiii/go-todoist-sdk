package go_todoist_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	client       *http.Client
	clientId     string
	clientSecret string
}

type Due struct {
	DueString string `json:"string"`
}

type Task struct {
	Id          string `json:"id"`
	Name        string `json:"content"`
	Priority    int    `json:"priority"`
	Description string `json:"description"`
	DueDate     Due    `json:"due"`
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
	date,
	description,
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

	requestURL := fmt.Sprintf("%s/rest/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT)
	requestBody := []byte(fmt.Sprintf(`{"content": "%s", 
										       "description": "%s",
											   "priority": %s,
											   "due_string": "%s"}`, name, description, priority, date))
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

func (t *Client) getAllTasks(accessToken string) ([]Task, error) {
	response, err := t.createGetTasksRequest(accessToken)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task

	_ = json.Unmarshal(responseBody, &tasks)

	return tasks, nil
}

func (t *Client) createGetTasksRequest(accessToken string) (*http.Response, error) {
	client := t.client
	requestURl := fmt.Sprintf("%s/rest/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT)
	req, err := http.NewRequest(http.MethodGet, requestURl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}

func (t *Client) createGetTaskRequest(taskID, accessToken string) (*http.Response, error) {
	client := t.client
	requestURL := fmt.Sprintf("%s/rest/%s/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT, taskID)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}

func (t *Client) GetTask(taskID, accessToken string) (*Task, error) {
	task := &Task{}
	response, err := t.createGetTaskRequest(taskID, accessToken)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal(responseBody, task)

	return task, nil
}
