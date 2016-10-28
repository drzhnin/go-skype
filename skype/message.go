package skype

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MessegesService service

type Messege struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (s *MessegesService) Send(conversationID string, messegeType string, text string) (*Response, error) {
	apiURL := fmt.Sprintf("https://api.skype.net/v3/conversations/%s/activities", conversationID)
	messege := Messege{Type: messegeType, Text: text}
	res, err := json.Marshal(messege)
	if err != nil {
		return nil, err
	}
	body := strings.NewReader(string(res))
	req, err := s.client.NewRequest("POST", apiURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	return resp, nil
}
