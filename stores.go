package rawg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetStores returns a list of video game storefronts
func (api *Client) GetStores(page int, pageSize int, ordering string) ([]*Store, int, error) {
	path := "/stores"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	if ordering != "" {
		data["ordering"] = ordering
	}

	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Store `json:"results"`
		Count   int      `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetStore returns details of the store
func (api *Client) GetStore(id int) (*StoreDetailed, error) {
	path := fmt.Sprintf("/stores/%d", id)
	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var store StoreDetailed

	if err := json.Unmarshal(body, &store); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &store, nil
}
