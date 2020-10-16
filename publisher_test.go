package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetPublisher(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	publisherID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/publisher.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/publishers/%d?lang=ru", apiBaseUrl, publisherID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	publisher, err := client.GetPublisher(publisherID)
	assert.NoError(t, err)
	assert.Equal(t, "Electronic Arts", publisher.Name)
}

func TestClient_GetPublisher_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	publisherID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/publishers/%d?lang=ru", apiBaseUrl, publisherID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	publisher, err := client.GetPublisher(publisherID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, publisher)
}

func TestClient_GetPublisher_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	publisherID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/publishers/%d?lang=ru", apiBaseUrl, publisherID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	publisher, err := client.GetPublisher(publisherID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, publisher)
}
