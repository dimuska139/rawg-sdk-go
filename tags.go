package rawg

import (
	"fmt"
)

// GetTags returns a list of tags
func (api *Client) GetTags(page int, pageSize int) ([]*Tag, int, error) {
	path := "/tags"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Tag `json:"results"`
		Count   int    `json:"count"`
	}

	if err := api.get(path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetTag returns details of the tag
func (api *Client) GetTag(id int) (*TagDetailed, error) {
	path := fmt.Sprintf("/tags/%d", id)
	var tag TagDetailed

	if err := api.get(path, nil, &tag); err != nil {
		return nil, err
	}

	return &tag, nil
}
