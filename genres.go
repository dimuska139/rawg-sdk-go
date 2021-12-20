package rawg

import (
	"context"
	"fmt"
)

// GetGenres returns a list of video game genres
func (api *Client) GetGenres(ctx context.Context, page int, pageSize int, ordering string) ([]*Genre, int, error) {
	path := "/genres"
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	if ordering != "" {
		data["ordering"] = ordering
	}

	var response struct {
		Results []*Genre `json:"results"`
		Count   int      `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGenre returns details of the genre
func (api *Client) GetGenre(ctx context.Context, id int) (*GenreDetailed, error) {
	path := fmt.Sprintf("/genres/%d", id)
	var genre GenreDetailed

	if err := api.get(ctx, path, nil, &genre); err != nil {
		return nil, err
	}

	return &genre, nil
}
