package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetGame(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/game.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d?lang=ru", apiBaseUrl, gameID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	game, err := client.GetGame(gameID)
	assert.NoError(t, err)
	assert.Equal(t, "Full Throttle Remastered", game.Name)
}

func TestClient_GetGame_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d?lang=ru", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	game, err := client.GetGame(gameID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, game)
}

func TestClient_GetGame_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d?lang=ru", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	game, err := client.GetGame(gameID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, game)
}
