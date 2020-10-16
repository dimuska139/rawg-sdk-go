package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetPublishers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/publishers.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/publishers?lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	publishers, total, err := client.GetPublishers(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(publishers))
	assert.Equal(t, "Electronic Arts", publishers[0].Name)
	assert.Equal(t, 44546, total)
}

func TestClient_GetPublishers_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/publishers?lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	publishers, total, err := client.GetPublishers(1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(publishers))
	assert.Equal(t, 0, total)
}

func TestClient_GetPublishers_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/publishers?lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	publishers, total, err := client.GetPublishers(1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(publishers))
	assert.Equal(t, 0, total)
}
