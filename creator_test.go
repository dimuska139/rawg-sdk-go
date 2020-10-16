package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetCreator(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creatorID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/creator.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/creators/%d?lang=ru", apiBaseUrl, creatorID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	creator, err := client.GetCreator(creatorID)
	assert.NoError(t, err)
	assert.Equal(t, "Michael Unsworth", creator.Name)
}

func TestClient_GetCreator_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creatorID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/creators/%d?lang=ru", apiBaseUrl, creatorID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	creator, err := client.GetCreator(creatorID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, creator)
}

func TestClient_GetCreator_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creatorID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/creators/%d?lang=ru", apiBaseUrl, creatorID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	creator, err := client.GetCreator(creatorID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, creator)
}
