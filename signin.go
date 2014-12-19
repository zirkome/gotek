package gotech

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
)

func SignIn(login string, password string) (*Board, error) {
	response, err := http.PostForm(API_ENDPOINT + "?format=json",
		url.Values{"login": {login}, "password": {password}})
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(contents[31:], &Dashboard)
		if err != nil {
			return nil, err
		}
	}
	return &Dashboard, nil
}
