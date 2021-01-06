package rawg

import (
	"fmt"
)

// GetPlatforms returns a list of video game platforms
func (api *Client) GetPlatforms(page int, pageSize int, ordering string) ([]*Platform, int, error) {
	path := "/platforms"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	if ordering != "" {
		data["ordering"] = ordering
	}

	var response struct {
		Results []*Platform `json:"results"`
		Count   int         `json:"count"`
	}

	if err := api.get(path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetParentsPlatforms returns a list of parent platforms
func (api *Client) GetParentsPlatforms(page int, pageSize int, ordering string) ([]*Platform, int, error) {
	path := "/platforms/lists/parents"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	if ordering != "" {
		data["ordering"] = ordering
	}

	var response struct {
		Results []*Platform `json:"results"`
		Count   int         `json:"count"`
	}

	if err := api.get(path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetPlatform returns details of the platform
func (api *Client) GetPlatform(id int) (*PlatformDetailed, error) {
	path := fmt.Sprintf("/platforms/%d", id)

	var platform PlatformDetailed

	if err := api.get(path, nil, &platform); err != nil {
		return nil, err
	}

	return &platform, nil
}
