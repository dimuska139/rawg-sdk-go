package rawg

import (
	"context"
	"fmt"
)

// GetDevelopers returns a list of game developers
func (api *Client) GetDevelopers(ctx context.Context, page int, pageSize int) ([]*Developer, int, error) {
	path := "/developers"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Developer `json:"results"`
		Count   int          `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetDeveloper returns details of the developer
func (api *Client) GetDeveloper(ctx context.Context, id int) (*DeveloperDetailed, error) {
	path := fmt.Sprintf("/developers/%d", id)
	var developer DeveloperDetailed
	if err := api.get(ctx, path, nil, &developer); err != nil {
		return nil, err
	}

	return &developer, nil
}
