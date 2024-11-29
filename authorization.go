package github.com/hhhanafubukiii/go-todoist-sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	authorizationURL, err := t.getUrl(AUTHORIZE_ENDPOINT)
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

	accessToken, err := t.doHTTP(
		TOKEN_ENDPOINT,
		code,
		tokenRequest.clientId,
		tokenRequest.clientSecret,
		requestBody,
	)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func (t *Client) doHTTP(path, code, clientID, clientSecret string, reqBody []byte) (string, error) {
	client := t.client
	tokenResponse := TokenResponse{}

	bodyReader := bytes.NewReader(reqBody)
	uri, _ := t.getUrl(path)
	requestURL := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", uri, clientID, clientSecret, code)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatal(err)
	}

	return tokenResponse.Token, nil
}

func (t *Client) getUrl(path string) (string, error) {
	return fmt.Sprintf("%s/%s", AUTH_BASE_URL, path), nil
}

func (t *TokenRequest) setBody() ([]byte, error) {
	s := fmt.Sprintf(`{
							"client_id": %s
							"client_secret": %s
							"code": %s
							}`, t.clientId, t.clientSecret, t.code)

	body, err := json.Marshal(s)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
