package liveagent

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type StringMap map[string]string

type Client struct {
	APIURL string
	APIKey string
}

func NewClient(url, key string) *Client {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	return &Client{url, key}
}

type ErrorResponse struct {
	Response struct {
		Status       string `json:"status"`
		Statuscode   int    `json:"statuscode"`
		Errormessage string `json:"errormessage"`
		Debugmessage string `json:"debugmessage"`
	} `json:"response"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s %d: %s", e.Response.Status, e.Response.Statuscode, e.Response.Errormessage)
}

func (c *Client) get(path string, params StringMap, out interface{}) error {
	var err error
	req, err := http.NewRequest("GET", c.APIURL+path, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("apikey", c.APIKey)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	client := c.httpclient()
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	/*
		decoder := json.NewDecoder(res.Body)
		if res.StatusCode == 200 {
			err = decoder.Decode(out)
			if err != nil {
				return err
			}
		} else {
			var laerr ErrorResponse
			err = decoder.Decode(&laerr)
			if err != nil {
				return err
			}
			return &laerr
		}

		return nil
	*/
	return c.handleResult(res, out)
}

func (c *Client) httpclient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}

func (c *Client) post(path string, in, out interface{}) error {
	var err error
	form, err := query.Values(in)
	if err != nil {
		return err
	}

	form["apikey"] = []string{c.APIKey}
	req, err := http.NewRequest("POST", c.APIURL+path, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := c.httpclient()
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return c.handleResult(res, out)
}

func (c *Client) handleResult(res *http.Response, out interface{}) error {
	var err error
	decoder := json.NewDecoder(res.Body)
	if res.StatusCode == 200 {
		err = decoder.Decode(out)
		if err != nil {
			return err
		}
	} else {
		var laerr ErrorResponse
		err = decoder.Decode(&laerr)
		if err != nil {
			return err
		}
		return &laerr
	}
	return nil
}
