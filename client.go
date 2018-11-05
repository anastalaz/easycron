package easycron

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client implements an Easycron client.
type Client struct {
	*Config
}

// New client.
func New(config *Config) *Client {
	c := &Client{Config: config}
	return c
}

// call the endpoint.
func (c *Client) call(method string, v url.Values) ([]byte, error) {
	url := c.Config.BaseURL + method

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = v.Encode()
	req.Header.Add("content-type", "application/json")

	r, _, err := c.do(req)
	return r, err
}

// do performs the request and handles easycron error.
func (c *Client) do(req *http.Request) ([]byte, int64, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	e := &Error{}

	if err := json.Unmarshal(body, e); err != nil {
		return nil, 0, err
	}

	if e.Status == "error" {
		return nil, 0, e
	}

	return body, res.ContentLength, nil
}
