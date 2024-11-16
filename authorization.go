package go_todoist_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (t *Client) GetAuthenticationURL(clientId, scope, state string) (string, error) {
	if scope == "" || clientId == "" || state == "" {
		panic("Empty values!")
	}

	authRequest := &AuthorizationRequest{
		clientId: clientId,
		scope:    scope,
		state:    state,
	}

	authorizationURL, err := getUrl(AUTHORIZE_ENDPOINT)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s?client_id=%s&scope=%s&state=%s", authorizationURL,
		authRequest.clientId,
		authRequest.scope,
		authRequest.state), nil
}

func (t *Client) GetAccessToken(clientId, clientSecret, code string) (string, error) {
	if clientSecret == "" || clientId == "" || code == "" {
		panic("Empty values!")
	}

	tokenRequest := &TokenRequest{
		clientId:     clientId,
		clientSecret: clientSecret,
		code:         code,
	}

	requestBody, err := tokenRequest.setBody()
	if err != nil {
		return "", err
	}
	accessToken, err := t.doHTTP(TOKEN_ENDPOINT, requestBody)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func getUrl(path string) (string, error) {
	return fmt.Sprintf("%s/%s", AUTH_BASE_URL, path), nil
}

func (t *TokenRequest) setBody() ([]byte, error) {
	s := fmt.Sprintf(`{
		"client_id": %s,
		"client_secret": %s,
		"code": %s
	}`, t.clientId, t.clientSecret, t.code)

	body, err := json.Marshal(s)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (t *Client) doHTTP(path string, body []byte) (string, error) {
	uri, err := getUrl(path)
	if err != nil {
		return "80", err
	}

	tokenReq := &TokenRequest{}

	link := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", uri, tokenReq.clientId, tokenReq.clientSecret, tokenReq.code)

	req, err := http.NewRequest(http.MethodPost, link, bytes.NewBuffer(body))
	if err != nil {
		return link, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := t.client.Do(req)
	if err != nil {
		return "93", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "96", errors.New(response.Status)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "101", err
	}

	values, err := url.ParseQuery(string(responseBody))
	if err != nil {
		return "106", err
	}

	return values.Get("access_token"), nil
}
