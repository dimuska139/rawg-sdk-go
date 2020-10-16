package rawg_sdk_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Client) GetPlatforms(page int, pageSize int, ordering string) ([]*Platform, int, error) {
	path := "/platforms"
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
		Results []*Platform `json:"results"`
		Count   int         `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

func (api *Client) GetParentsPlatforms(page int, pageSize int, ordering string) ([]*Platform, int, error) {
	path := "/platforms/lists/parents"
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
		Results []*Platform `json:"results"`
		Count   int         `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

func (api *Client) GetPlatform(id int) (*PlatformDetailed, error) {
	path := fmt.Sprintf("/platforms/%d", id)
	body, err := api.NewRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var platform PlatformDetailed

	if err := json.Unmarshal(body, &platform); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &platform, nil
}
