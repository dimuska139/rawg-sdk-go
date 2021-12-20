package rawg

import (
	"context"
	"fmt"
)

// GetStores returns a list of video game storefronts
func (api *Client) GetStores(ctx context.Context, page int, pageSize int, ordering string) ([]*Store, int, error) {
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

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetStore returns details of the store
func (api *Client) GetStore(ctx context.Context, id int) (*StoreDetailed, error) {
	path := fmt.Sprintf("/stores/%d", id)
	var store StoreDetailed

	if err := api.get(ctx, path, nil, &store); err != nil {
		return nil, err
	}

	return &store, nil
}
