package rawg

import (
	"context"
	"fmt"
)

// GetCreatorRoles returns a list of creator positions (jobs)
func (api *Client) GetCreatorRoles(ctx context.Context, page int, pageSize int) ([]*Role, int, error) {
	path := "/creator-roles"
	params := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Role `json:"results"`
		Count   int     `json:"count"`
	}

	if err := api.get(ctx, path, params, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetCreators returns a list of game creators
func (api *Client) GetCreators(ctx context.Context, page int, pageSize int) ([]*Creator, int, error) {
	path := "/creators"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Creator `json:"results"`
		Count   int        `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetCreator returns details of the creator
func (api *Client) GetCreator(ctx context.Context, id int) (*CreatorDetailed, error) {
	path := fmt.Sprintf("/creators/%d", id)
	var creator CreatorDetailed

	if err := api.get(ctx, path, nil, &creator); err != nil {
		return nil, err
	}

	return &creator, nil
}
