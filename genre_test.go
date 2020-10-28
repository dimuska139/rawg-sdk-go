package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetGenre(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	genreID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/genre.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/genres/%d?key=test&lang=ru", apiBaseUrl, genreID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	genre, err := client.GetGenre(genreID)
	assert.NoError(t, err)
	assert.Equal(t, "Racing", genre.Name)
}

func TestClient_GetGenre_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	genreID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/genres/%d?key=test&lang=ru", apiBaseUrl, genreID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	genre, err := client.GetGenre(genreID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, genre)
}

func TestClient_GetGenre_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	genreID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/genres/%d?key=test&lang=ru", apiBaseUrl, genreID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	genre, err := client.GetGenre(genreID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, genre)
}
