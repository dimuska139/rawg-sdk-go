package rawg_sdk_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Client) GetPublishers(page int, pageSize int) ([]*Publisher, int, error) {
	path := "/publishers"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.NewRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Publisher `json:"results"`
		Count   int          `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

func (api *Client) GetPublisher(id int) (*PublisherDetailed, error) {
	path := fmt.Sprintf("/publishers/%d", id)
	body, err := api.NewRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var publisher PublisherDetailed

	if err := json.Unmarshal(body, &publisher); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &publisher, nil
}