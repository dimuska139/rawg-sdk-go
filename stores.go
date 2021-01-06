package rawg

import (
	"fmt"
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

	var response struct {
		Results []*Store `json:"results"`
		Count   int      `json:"count"`
	}

	if err := api.get(path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetStore returns details of the store
func (api *Client) GetStore(id int) (*StoreDetailed, error) {
	path := fmt.Sprintf("/stores/%d", id)
	var store StoreDetailed

	if err := api.get(path, nil, &store); err != nil {
		return nil, err
	}

	return &store, nil
}
