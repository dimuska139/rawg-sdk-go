package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetStores(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/stores.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/stores?key=test&lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	stores, total, err := client.GetStores(1, 2, "-name")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(stores))
	assert.Equal(t, "Steam", stores[0].Name)
	assert.Equal(t, 10, total)
}

func TestClient_GetStores_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/stores?key=test&lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	stores, total, err := client.GetStores(1, 2, "-name")
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(stores))
	assert.Equal(t, 0, total)
}

func TestClient_GetStores_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/stores?key=test&lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	stores, total, err := client.GetStores(1, 2, "-name")
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(stores))
	assert.Equal(t, 0, total)
}
