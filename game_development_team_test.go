package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetGameDevelopmentTeam(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/development_team.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/development-team?lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl, gameID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, total, err := client.GetGameDevelopmentTeam(gameID, 1, 2, "-name")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(items))
	assert.Equal(t, 5, total)
	assert.Equal(t, "Keisuke Kikuchi", items[0].Name)
}

func TestClient_GetGameDevelopmentTeam_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/development-team?lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, total, err := client.GetGameDevelopmentTeam(gameID, 1, 2, "-name")
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(items))
	assert.Equal(t, 0, total)
}

func TestClient_GetGameDevelopmentTeam_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gameID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/games/%d/development-team?lang=ru&ordering=-name&page=1&page_size=2", apiBaseUrl, gameID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	items, total, err := client.GetGameDevelopmentTeam(gameID, 1, 2, "-name")
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(items))
	assert.Equal(t, 0, total)
}
