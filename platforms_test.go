package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetPlatforms(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/platforms.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/platforms?lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	platforms, total, err := client.GetPlatforms(1, 2, "-name")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(platforms))
	assert.Equal(t, "PC", platforms[0].Name)
	assert.Equal(t, 51, total)
}

func TestClient_GetPlatforms_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/platforms?lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	platforms, total, err := client.GetPlatforms(1, 2, "-name")
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(platforms))
	assert.Equal(t, 0, total)
}

func TestClient_GetPlatforms_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/platforms?lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	platforms, total, err := client.GetPlatforms(1, 2, "-name")
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(platforms))
	assert.Equal(t, 0, total)
}
