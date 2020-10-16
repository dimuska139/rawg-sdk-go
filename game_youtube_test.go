package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetGameYoutube(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/game_youtube.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/youtube?lang=ru", apiBaseUrl, gameID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, err := client.GetGameYoutube(gameID)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(items))
	assert.Equal(t, "ENTER THE LOST VILLAGE | Fatal Frame 2 - Part 1", items[0].Name)
}

func TestClient_GetGameYoutube_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/youtube?lang=ru", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, err := client.GetGameYoutube(gameID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(items))
}

func TestClient_GetGameYoutube_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/youtube?lang=ru", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, err := client.GetGameYoutube(gameID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(items))
}
