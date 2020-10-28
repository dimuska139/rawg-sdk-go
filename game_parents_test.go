package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetGameParents(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/parent_games.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/parent-games?key=test&lang=ru&page=1&page_size=2", apiBaseUrl, gameID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, total, err := client.GetParentGames(gameID, 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, 1, total)
	assert.Equal(t, "Crash Bandicoot 3: Warped", items[0].Name)
}

func TestClient_GetGameParents_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/parent-games?key=test&lang=ru&page=1&page_size=2", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, total, err := client.GetParentGames(gameID, 1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(items))
	assert.Equal(t, 0, total)
}

func TestClient_GetGameParents_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/parent-games?key=test&lang=ru&page=1&page_size=2", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, total, err := client.GetParentGames(gameID, 1, 2)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(items))
	assert.Equal(t, 0, total)
}
