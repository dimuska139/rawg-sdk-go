package rawgSdkGo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get a list of game developers
func (api *Client) GetDevelopers(page int, pageSize int) ([]*Developer, int, error) {
	path := "/developers"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Developer `json:"results"`
		Count   int          `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// Get details of the developer
func (api *Client) GetDeveloper(id int) (*DeveloperDetailed, error) {
	path := fmt.Sprintf("/developers/%d", id)
	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var developer DeveloperDetailed

	if err := json.Unmarshal(body, &developer); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &developer, nil
}
