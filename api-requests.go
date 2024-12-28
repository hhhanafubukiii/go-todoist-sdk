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

type Task struct {
	Id          string `json:"id"`
	Name        string `json:"content"`
	Priority    int    `json:"priority"`
	Description string `json:"description"`
	DueDate     Due    `json:"due"`
}

type AddTask struct {
	Name        string
	Priority    string
	DueDate     string
	Description string
}

type AuthorizationRequest struct {
	clientId string
	scope    string
	state    string
}

type AuthorizationResponse struct {
	code  string
	state string
}

type TokenRequest struct {
	clientId     string
	clientSecret string
	code         string
}

type TokenResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

type Due struct {
	DueString string `json:"string"`
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

func (t *Client) CloseTask(taskID, accessToken string) error {
	requestURL := fmt.Sprintf("%s/rest/%s/%s/%s/close", BASE_URL, REST_VERSION, TASKS_ENDPOINT, taskID)
	req, err := http.NewRequest(http.MethodPost, requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	response, err := t.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

func (t *Client) DeleteTask(taskID, accessToken string) error {
	requestURL := fmt.Sprintf("%s/rest/%s/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT, taskID)
	req, err := http.NewRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	response, err := t.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

func (t *Client) UpdateTaskName(taskID, accessToken, newName string) error {
	requestURL := fmt.Sprintf("%s/rest/%s/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT, taskID)
	requestBody := []byte(fmt.Sprintf(`{
												"content": "%s"
										      }`, newName))
	bodyReader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")

	response, err := t.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	return nil
}

func (t *Client) UpdateTaskPriority(taskID, accessToken, newPriority string) error {
	requestURL := fmt.Sprintf("%s/rest/%s/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT, taskID)
	requestBody := []byte(fmt.Sprintf(`{
												"priority": "%s"
										      }`, newPriority))
	bodyReader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")

	response, err := t.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	return nil
}

func (t *Client) UpdateTaskDueDate(taskID, accessToken, newDueDate string) error {
	requestURL := fmt.Sprintf("%s/rest/%s/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT, taskID)
	requestBody := []byte(fmt.Sprintf(`{
												"due_string": "%s"
										      }`, newDueDate))
	bodyReader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")

	response, err := t.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	return nil
}

func (t *Client) UpdateTaskDescription(taskID, accessToken, newDescription string) error {
	requestURL := fmt.Sprintf("%s/rest/%s/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT, taskID)
	requestBody := []byte(fmt.Sprintf(`{
												"description": "%s"
										      }`, newDescription))
	bodyReader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")

	response, err := t.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	return nil
}

func (t *Client) GetAllTasks(accessToken string) ([]Task, error) {
	response, err := t.createGetTasksRequest(accessToken, "")
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

func (t *Client) GetTodayTasks(accessToken string) ([]Task, error) {
	response, err := t.createGetTasksRequest(accessToken, "today")
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

func (t *Client) createGetTasksRequest(accessToken, filter string) (*http.Response, error) {
	client := t.client
	requestURl := fmt.Sprintf("%s/rest/%s/%s", BASE_URL, REST_VERSION, TASKS_ENDPOINT)
	if filter != "" {
		requestURl = fmt.Sprintf("%s?filter=%s", requestURl, filter)
	}
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
