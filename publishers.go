package rawg

import (
	"context"
	"fmt"
)

// GetPublishers returns a list of video game publishers
func (api *Client) GetPublishers(ctx context.Context, page int, pageSize int) ([]*Publisher, int, error) {
	path := "/publishers"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Publisher `json:"results"`
		Count   int          `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetPublisher returns details of the publisher
func (api *Client) GetPublisher(ctx context.Context, id int) (*PublisherDetailed, error) {
	path := fmt.Sprintf("/publishers/%d", id)
	var publisher PublisherDetailed

	if err := api.get(ctx, path, nil, &publisher); err != nil {
		return nil, err
	}

	return &publisher, nil
}
