package gotodoistsdk

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func (t *AuthorizationRequest) GetAuthenticationURL(client_id, scope, state string) (string, error) {
	if scope == "" || client_id == "" || state == "" {
		log.Fatal("Empty params")
	}

	t.client_id = client_id
	t.scope = scope
	t.state = state

	authorizationURL := getUrl(AUTHORIZE_ENDPOINT)

	return fmt.Sprintf("%s?client_id=%s&scope=%s&state=%s", authorizationURL, t.client_id, t.scope, t.state), nil
}

func (t *TokenRequest) GetAccessToken(client_id, client_secret, code string) (*TokenResponse, error) {
	if client_secret == "" || client_id == "" || code == "" {
		panic("Empty values!")
	}

	t.client_id = client_id
	t.client_secret = client_secret
	t.code = code

	url := getUrl(TOKEN_ENDPOINT)

	body, err := t.getBody()

	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		panic(err)
	}

	response, err := t.sendRequest(request)

	if err != nil {
		panic(err)
	}

	// return ...
}

func getUrl(path string) string {
	return fmt.Sprintf("%s/%s", AUTH_BASE_URL, path)
}

func (t *TokenRequest) getBody() ([]byte, error) {
	s := fmt.Sprintf(`{
		"client_id": %s,
		"client_secret": %s,
		"code": %s
	}`, t.client_id, t.client_secret, t.code)

	return []byte(s), nil
}

func (t *TokenRequest) sendRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	return response, nil
}
