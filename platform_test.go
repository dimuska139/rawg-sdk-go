package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetPlatform(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	platformID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/platform.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/platforms/%d?lang=ru", apiBaseUrl, platformID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	platform, err := client.GetPlatform(platformID)
	assert.NoError(t, err)
	assert.Equal(t, "Xbox One", platform.Name)
}

func TestClient_GetPlatform_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	platformID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/platforms/%d?lang=ru", apiBaseUrl, platformID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	platform, err := client.GetPlatform(platformID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, platform)
}

func TestClient_GetPlatform_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	platformID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/platforms/%d?lang=ru", apiBaseUrl, platformID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	platform, err := client.GetPlatform(platformID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, platform)
}
