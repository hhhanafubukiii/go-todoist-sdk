package gotodoistsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *AuthorizationRequest) GetAuthenticationURL(client_id, scope, state string) (string, error) {
	if scope == "" || client_id == "" || state == "" {
		panic("Empty values!")
	}

	authRequest := &AuthorizationRequest{
		client_id: client_id,
		scope:     scope,
		state:     state,
	}

	authorizationURL, err := getUrl(AUTHORIZE_ENDPOINT)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s?client_id=%s&scope=%s&state=%s", authorizationURL,
		authRequest.client_id,
		authRequest.scope,
		authRequest.state), nil
}

func (t *TokenRequest) GetAccessToken(client_id, client_secret, code string) (string, error) {
	if client_secret == "" || client_id == "" || code == "" {
		panic("Empty values!")
	}

	tokenRequest := &TokenRequest{
		client_id:     client_id,
		client_secret: client_secret,
		code:          code,
	}

	url, err := getUrl(TOKEN_ENDPOINT)
	if err != nil {
		panic(err)
	}

	body, err := t.setBody()
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	response, err := t.sendRequest(request)
	if err != nil {
		panic(err)
	}

	TokenResponse := &TokenResponse{}
	dr := json.NewDecoder(response.Body).Decode(TokenResponse)
	if dr != nil {
		panic(dr)
	}

	if response.StatusCode != 200 {
		panic(response.Status)
	}

	return TokenResponse.Token, nil
}

func getUrl(path string) (string, error) {
	return fmt.Sprintf("%s/%s", AUTH_BASE_URL, path), nil
}

func (t *TokenRequest) setBody() ([]byte, error) {
	s := fmt.Sprintf(`{
		"client_id": %s,
		"client_secret": %s,
		"code": %s
	}`, t.client_id, t.client_secret, t.code)

	return []byte(s), nil
}

func (t *TokenRequest) sendRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	defer request.Body.Close()

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response, nil
}
