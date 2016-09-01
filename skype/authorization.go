package skype

import (
	"fmt"
	"strings"
)

type AuthorizeService service

const (
	authURL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
)

type Authorization struct {
	TokenType    string `json:"token_type, omitempty"`
	ExpiresIn    int    `json:"expires_in, omitempty"`
	ExtExpiresIn int    `json:"ext_expires_in, omitempty"`
	AccessToken  string `json:"access_token, omitempty"`
}

func (s *AuthorizeService) Authorize() (*Authorization, *Response, error) {
	bodyStr := fmt.Sprintf("client_id=%s&scope=https://graph.microsoft.com/.default&grant_type=client_credentials&client_secret=%s", s.client.ClientID, s.client.ClientSecret)
	body := strings.NewReader(bodyStr)
	req, err := s.client.NewRequest("POST", authURL, body)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, nil, err
	}

	var auth = new(Authorization)

	resp, err := s.client.Do(req, auth)
	if err != nil {
		return nil, resp, err
	}

	return auth, resp, err
}
