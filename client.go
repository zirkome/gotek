package gotek

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client struct {
	Login    string
	Password string
}

var (
	client *http.Client
)

func init() {
	cookiejar, _ := cookiejar.New(nil)
	client = &http.Client{Jar: cookiejar}
}

func (c *Client) SignIn() (*Dashboard, error) {
	var dashboard Dashboard

	response, err := client.PostForm(DASHBOARD_ENDPOINT,
		url.Values{
			"login":    {c.Login},
			"password": {c.Password},
		})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &dashboard)
	if err != nil {
		return nil, err
	}
	return &dashboard, nil
}
