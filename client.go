package rawg

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"io/ioutil"
	"net/http"
	"regexp"
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
	baseUrl     string
	client      *http.Client
	config      *Config
	rateLimiter *rate.Limiter
}

// NewClient creates new Client to interract with RAWG API
func NewClient(client *http.Client, config *Config) *Client {
	if config.Rps == 0 {
		config.Rps = 5
	}

	if config.Language == "" {
		config.Language = "en"
	}

	return &Client{
		baseUrl:     apiBaseUrl,
		client:      client,
		config:      config,
		rateLimiter: rate.NewLimiter(rate.Limit(config.Rps), config.Rps),
	}
}

func (api *Client) get(path string, params map[string]interface{}, responseModel interface{}) error {
	fullPath := api.baseUrl + path

	ctx := context.Background()
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return err
	}

	req, e := http.NewRequest(http.MethodGet, fullPath, nil)
	if e != nil {
		return e
	}
	req.Header.Add("content-type", "application/json;charset=utf-8")
	q := req.URL.Query()

	// Workaround to switch to using API key (not all methods of the RAWG API support it)
	suggestedURLr, _ := regexp.Compile("/games/[0-9+]/suggested")
	twitchURLr, _ := regexp.Compile("/games/[0-9+]/twitch")
	youtubeURLr, _ := regexp.Compile("/games/[0-9+]/youtube")
	if !suggestedURLr.MatchString(path) && !twitchURLr.MatchString(path) && !youtubeURLr.MatchString(path) {
		q.Add("key", api.config.ApiKey)
	}
	//////////////////////////////////////////////////////////////////////////////////////

	q.Add("lang", api.config.Language)
	for param, value := range params {
		q.Add(param, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = q.Encode()
	resp, err := api.client.Do(req)
	if err != nil {
		return &RawgError{http.StatusServiceUnavailable, path, "", err.Error()}
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &RawgError{resp.StatusCode, path, string(body), err.Error()}
	}

	if resp.StatusCode != http.StatusOK {
		return &RawgError{resp.StatusCode, path, string(body), ""}
	}

	if err := json.Unmarshal(body, responseModel); err != nil {
		return fmt.Errorf("could not decode the data: %s", err)
	}

	return nil
}
