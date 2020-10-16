package rawg

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const apiBaseUrl = "https://api.rawg.io/api"

// RawgError represents information about errors
type RawgError struct {
	HttpCode int    // HTTP status code (https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
	Url      string // URL of RAWG endpoint associated with callable function
	Body     string // Raw body of response
	Message  string // Any comment
}

// Error() converts error to string
func (e *RawgError) Error() string {
	return fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Message)
}

// Client to interract with RAWG API
type Client struct {
	client      *http.Client
	config      *Config
	rateLimiter *rate.Limiter
}

// NewClient creates new Client to interract with RAWG API
func NewClient(client *http.Client, config *Config) *Client {
	rps := config.Rps
	if rps == 0 {
		rps = 5
	}

	return &Client{
		client:      client,
		config:      config,
		rateLimiter: rate.NewLimiter(rate.Every(1*time.Second), rps),
	}
}

// newRequest: Makes request with appending required params
func (api *Client) newRequest(path string, method string, data map[string]interface{}) ([]byte, error) {
	q := url.Values{}
	for param, value := range data {
		q.Add(param, fmt.Sprintf("%v", value))
	}

	q.Add("lang", api.config.Language)

	method = strings.ToUpper(method)
	fullPath := apiBaseUrl + path

	ctx := context.Background()
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	req, e := http.NewRequest(method, fullPath, bytes.NewBufferString(q.Encode()))
	if e != nil {
		return nil, e
	}

	if method == http.MethodGet {
		req.URL.RawQuery = q.Encode()
		req.Body = nil
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	}

	req.Header.Add("User-Agent", api.config.AppName)

	resp, err := api.client.Do(req)
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
