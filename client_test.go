package rawg

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"net/http"
)

func (suite *RAWGTestSuite) TestNewClient() {
	config := Config{
		ApiKey:   "anykey",
		Language: "en",
		Rps:      3,
	}
	client := NewClient(http.DefaultClient, &config)
	suite.Equal("https://api.rawg.io/api", client.baseUrl)
	suite.Equal(config.ApiKey, client.config.ApiKey)
	suite.Equal(config.Rps, client.config.Rps)
}

func (suite *RAWGTestSuite) TestNewClientWithoutDefaultParams() {
	config := Config{
		ApiKey: "anykey",
	}
	client := NewClient(http.DefaultClient, &config)
	suite.Equal(5, client.config.Rps)
	suite.Equal("en", client.config.Language)
}

func (suite *RAWGTestSuite) TestHttpError() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	badStatusCode := http.StatusInternalServerError

	config := Config{
		Language: "ru",
		ApiKey:   "anykey",
	}
	client := NewClient(http.DefaultClient, &config)
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/test?key=%s&lang=%s", apiBaseUrl, config.ApiKey, config.Language),
		httpmock.NewStringResponder(badStatusCode, ""),
	)

	err := client.get(context.Background(), "/test", nil, nil)
	suite.Error(err)
	_, isResponseError := err.(*RawgError)
	suite.True(isResponseError)
	suite.Equal(badStatusCode, err.(*RawgError).HttpCode)
}

func (suite *RAWGTestSuite) TestInvalidJsonError() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := Config{
		Language: "ru",
		ApiKey:   "anykey",
	}
	client := NewClient(http.DefaultClient, &config)
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/test?key=%s&lang=%s", apiBaseUrl, config.ApiKey, config.Language),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	err := client.get(context.Background(), "/test", nil, nil)
	suite.Error(err)
	_, isResponseError := err.(*RawgError)
	suite.False(isResponseError)
}

func (suite *RAWGTestSuite) TestRawgError() {
	err := RawgError{
		HttpCode: http.StatusInternalServerError,
		Url:      "/test",
		Body:     "Something went wrong",
		Message:  "Internal server error",
	}

	suite.Equal(fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", err.HttpCode, err.Url, err.Body, err.Message), err.Error())
}
