package rawg

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetTag(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tagID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/tag.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/tags/%d?key=test&lang=ru", apiBaseUrl, tagID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	tag, err := client.GetTag(tagID)
	assert.NoError(t, err)
	assert.Equal(t, "Survival", tag.Name)
}

func TestClient_GetTag_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tagID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/tags/%d?key=test&lang=ru", apiBaseUrl, tagID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	tag, err := client.GetTag(tagID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, tag)
}

func TestClient_GetTag_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tagID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/tags/%d?key=test&lang=ru", apiBaseUrl, tagID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := Config{
		ApiKey:   "test",
		Language: "ru",
		Rps:      5,
	}
	client := NewClient(http.DefaultClient, &config)
	tag, err := client.GetTag(tagID)
	assert.Error(t, err)
	_, isResponseError := err.(*RawgError)
	assert.True(t, isResponseError)
	assert.Nil(t, tag)
}
