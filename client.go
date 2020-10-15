package rawg_sdk_go

import (
	"bytes"
	"fmt"
	"golang.org/x/time/rate"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const apiBaseUrl = "https://api.rawg.io/api"

type RawgError struct {
	HttpCode int
	Url      string
	Body     string
	Message  string
}

func (e *RawgError) Error() string {
	return fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Message)
}

type Client struct {
	client      *http.Client
	config      *Config
	rateLimiter *rate.Limiter
}

func NewClient(client *http.Client, config *Config, rps int) *Client {
	if rps == 0 {
		rps = 5
	}

	return &Client{
		client:      client,
		config:      config,
		rateLimiter: rate.NewLimiter(rate.Every(1*time.Second), rps),
	}
}

func (c *Client) NewRequest(path string, method string, data map[string]interface{}) ([]byte, error) {
	q := url.Values{}
	for param, value := range data {
		q.Add(param, fmt.Sprintf("%v", value))
	}

	q.Add("lang", c.config.Language)

	method = strings.ToUpper(method)

	fullPath := apiBaseUrl + path
	req, e := http.NewRequest(method, fullPath, bytes.NewBufferString(q.Encode()))
	if e != nil {
		return nil, e
	}

	if method == "GET" {
		req.URL.RawQuery = q.Encode()
		req.Body = nil
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	}

	req.Header.Add("User-Agent", c.config.AppName)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &RawgError{http.StatusServiceUnavailable, path, "", err.Error()}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &RawgError{resp.StatusCode, path, string(body), err.Error()}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &RawgError{resp.StatusCode, path, string(body), ""}
	}

	return body, nil
}