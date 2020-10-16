package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetDeveloper(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	developerID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/developer.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/developers/%d?lang=ru", apiBaseUrl, developerID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	developer, err := client.GetDeveloper(developerID)
	assert.NoError(t, err)
	assert.Equal(t, "Feral Interactive", developer.Name)
}

func TestClient_GetDeveloper_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	developerID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/developers/%d?lang=ru", apiBaseUrl, developerID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	developer, err := client.GetDeveloper(developerID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, developer)
}

func TestClient_GetDeveloper_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	developerID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/developers/%d?lang=ru", apiBaseUrl, developerID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	developer, err := client.GetDeveloper(developerID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, developer)
}
