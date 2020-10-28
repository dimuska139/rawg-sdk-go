package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetCreators(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/creators.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/creators?key=test&lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	creators, total, err := client.GetCreators(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(creators))
	assert.Equal(t, "Cris Velasco", creators[0].Name)
	assert.Equal(t, 24406, total)
}

func TestClient_GetCreators_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/creators?key=test&lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	creators, total, err := client.GetCreators(1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(creators))
	assert.Equal(t, 0, total)
}

func TestClient_GetCreators_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/creators?key=test&lang=ru&page=1&page_size=2", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	creators, total, err := client.GetCreators(1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(creators))
	assert.Equal(t, 0, total)
}
