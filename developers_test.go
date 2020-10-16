package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetDevelopers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/developers.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/developers?lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	developers, total, err := client.GetDevelopers(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(developers))
	assert.Equal(t, "Feral Interactive", developers[0].Name)
	assert.Equal(t, 212621, total)
}

func TestClient_GetDevelopers_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/developers?lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	developers, total, err := client.GetDevelopers(1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(developers))
	assert.Equal(t, 0, total)
}

func TestClient_GetDevelopers_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/developers?lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	developers, total, err := client.GetDevelopers(1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(developers))
	assert.Equal(t, 0, total)
}
