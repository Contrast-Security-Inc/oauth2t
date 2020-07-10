package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

type Auth0TokenSource struct {
	issuer       string
	clientID     string
	clientSecret string
	audience     string
}

func NewAuth0TokenSource(issuer, clientID, clientSecret, audience string) *Auth0TokenSource {
	return &Auth0TokenSource{
		issuer:       issuer,
		clientID:     clientID,
		clientSecret: clientSecret,
		audience:     audience,
	}
}

func (ts *Auth0TokenSource) Token() (*oauth2.Token, error) {

	url := ts.issuer + "oauth/token"
	payload := strings.NewReader(
		fmt.Sprintf("grant_type=%s&client_id=%s&client_secret=%s&audience=%s",
			"client_credentials", ts.clientID, ts.clientSecret, ts.audience))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var token oauth2.Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
