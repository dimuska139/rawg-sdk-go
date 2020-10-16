package rawg_sdk_go

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetGames(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/games.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s/games?lang=ru&page=1&page_size=2&search=gta5`, apiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	/*dateRangeFirst := DateRange{
		From: time.Time{},
		To:   time.Time{},
	}

	dateRangeSecond := DateRange{
		From: time.Time{},
		To:   time.Time{},
	}*/
	filter := NewGamesFilter().
		SetPage(1).
		SetPageSize(2).
		SetSearch("gta5")
	/*.
	SetParentPlatforms(1, 2).
	SetPlatforms(3, 4).
	SetStores(5, 6).
	SetDevelopers(7, "feral-interactive").
	SetPublishers(8, "electronic-arts").
	SetGenres(9, "action", "indie").
	SetTags("singleplayer", 31).
	SetCreators(28, "mike-morasky").
	SetDates()*/

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	games, total, err := client.GetGames(filter)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(games))
	assert.Equal(t, "Grand Theft Auto V", games[0].Name)
	assert.Equal(t, 454984, total)
}

func TestClient_GetGames_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s/games?lang=ru&page=1&page_size=2&search=gta5`, apiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	filter := NewGamesFilter().
		SetPage(1).
		SetPageSize(2).
		SetSearch("gta5")

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	games, total, err := client.GetGames(filter)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(games))
	assert.Equal(t, 0, total)
}

func TestClient_GetGames_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s/games?lang=ru&page=1&page_size=2&search=gta5`, apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	filter := NewGamesFilter().
		SetPage(1).
		SetPageSize(2).
		SetSearch("gta5")

	config := Config{
		AppName:  "Test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	games, total, err := client.GetGames(filter)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)

	assert.Equal(t, 0, len(games))
	assert.Equal(t, 0, total)
}
