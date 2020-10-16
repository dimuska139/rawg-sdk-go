package rawg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTags returns a list of tags
func (api *Client) GetTags(page int, pageSize int) ([]*Tag, int, error) {
	path := "/tags"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Tag `json:"results"`
		Count   int    `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetTag returns details of the tag
func (api *Client) GetTag(id int) (*TagDetailed, error) {
	path := fmt.Sprintf("/tags/%d", id)
	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var tag TagDetailed

	if err := json.Unmarshal(body, &tag); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &tag, nil
}
