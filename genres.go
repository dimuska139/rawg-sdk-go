package rawg_sdk_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Client) GetGenres(page int, pageSize int, ordering string) ([]*Genre, int, error) {
	path := "/genres"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	if ordering != "" {
		data["ordering"] = ordering
	}

	body, err := api.NewRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Genre `json:"results"`
		Count   int      `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

func (api *Client) GetGenre(id int) (*GenreDetailed, error) {
	path := fmt.Sprintf("/genres/%d", id)
	body, err := api.NewRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var genre GenreDetailed

	if err := json.Unmarshal(body, &genre); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &genre, nil
}
