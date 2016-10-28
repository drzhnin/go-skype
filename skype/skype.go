package skype

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	client       *http.Client
	ClientID     string
	ClientSecret string
	Token        string

	Authorization *AuthorizeService
	Messege       *MessegesService
}

type service struct {
	client *Client
}

func NewClient(clientID, clientSecret string) *Client {
	httpClient := http.DefaultClient

	c := &Client{client: httpClient, ClientID: clientID, ClientSecret: clientSecret}
	c.Authorization = &AuthorizeService{client: c}
	c.Messege = &MessegesService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.Token != "" {
		req.Header.Add("authorization", c.Token)
	}

	return req, nil
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	response := newResponse(resp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}
	defer resp.Body.Close()

	return response, err
}
