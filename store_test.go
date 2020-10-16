package rawgSdkGo

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetStore(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	storeID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/store.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/stores/%d?lang=ru", apiBaseUrl, storeID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	store, err := client.GetStore(storeID)
	assert.NoError(t, err)
	assert.Equal(t, "Steam", store.Name)
}

func TestClient_GetStore_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	storeID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/stores/%d?lang=ru", apiBaseUrl, storeID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	store, err := client.GetStore(storeID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, store)
}

func TestClient_GetStore_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	storeID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/stores/%d?lang=ru", apiBaseUrl, storeID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	store, err := client.GetStore(storeID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, store)
}
