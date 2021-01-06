package rawg

import (
	"fmt"
)

// GetCreatorRoles returns a list of creator positions (jobs)
func (api *Client) GetCreatorRoles(page int, pageSize int) ([]*Role, int, error) {
	path := "/creator-roles"
	params := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Role `json:"results"`
		Count   int     `json:"count"`
	}

	if err := api.get(path, params, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetCreators returns a list of game creators
func (api *Client) GetCreators(page int, pageSize int) ([]*Creator, int, error) {
	path := "/creators"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Creator `json:"results"`
		Count   int        `json:"count"`
	}

	if err := api.get(path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetCreator returns details of the creator
func (api *Client) GetCreator(id int) (*CreatorDetailed, error) {
	path := fmt.Sprintf("/creators/%d", id)
	var creator CreatorDetailed

	if err := api.get(path, nil, &creator); err != nil {
		return nil, err
	}

	return &creator, nil
}
