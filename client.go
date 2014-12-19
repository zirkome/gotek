package gotek

import (
	"io/ioutil"
	"encoding/json"
	"net/url"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	Login		string
	Password	string
}

var (
	client *http.Client
)

func init() {
	cookiejar, _ := cookiejar.New(nil)
	client = &http.Client{ Jar: cookiejar }
}

func (c *Client) SignIn() (*Board, error) {
	var dashboard Board

	response, err := client.PostForm(DASHBOARD_ENDPOINT,
		url.Values{
			"login": { c.Login },
			"password": { c.Password },
		})
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(contents[31:], &dashboard)
		if err != nil {
			return nil, err
		}
	}
	return &dashboard, nil
}
