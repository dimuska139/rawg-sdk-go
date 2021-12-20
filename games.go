package rawg

import (
	"context"
	"fmt"
)

// GetGames returns a list of games
func (api *Client) GetGames(ctx context.Context, filter *GamesFilter) ([]*Game, int, error) {
	path := "/games"

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := api.get(ctx, path, filter.GetParams(), &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameAdditions returns a list of DLC's for the game, GOTY and other editions, companion apps, etc
func (api *Client) GetGameAdditions(ctx context.Context, gameID int, page int, pageSize int) ([]*Game, int, error) {
	path := fmt.Sprintf("/games/%d/additions", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameDevelopmentTeam returns a list of individual creators that were part of the development team
func (api *Client) GetGameDevelopmentTeam(ctx context.Context, gameID int, page int, pageSize int, ordering string) ([]*GameDeveloper, int, error) {
	path := fmt.Sprintf("/games/%d/development-team", gameID)
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	if ordering != "" {
		data["ordering"] = ordering
	}

	var response struct {
		Results []*GameDeveloper `json:"results"`
		Count   int              `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameSeries returns a list of games that are part of same series
func (api *Client) GetGameSeries(ctx context.Context, gameID int, page int, pageSize int) ([]*Game, int, error) {
	path := fmt.Sprintf("/games/%d/game-series", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetParentGames returns a list of parent games for DLC's and editions
func (api *Client) GetParentGames(ctx context.Context, gameID int, page int, pageSize int) ([]*Game, int, error) {
	path := fmt.Sprintf("/games/%d/parent-games", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameScreenshots returns a screenshots list for the game
func (api *Client) GetGameScreenshots(ctx context.Context, gameID int, page int, pageSize int) ([]*Screenshot, int, error) {
	path := fmt.Sprintf("/games/%d/screenshots", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*Screenshot `json:"results"`
		Count   int           `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameStores returns links to the stores that sell the game
func (api *Client) GetGameStores(ctx context.Context, gameID int, page int, pageSize int) ([]*GameStore, int, error) {
	path := fmt.Sprintf("/games/%d/stores", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	var response struct {
		Results []*GameStore `json:"results"`
		Count   int          `json:"count"`
	}

	if err := api.get(ctx, path, data, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGame returns details of the game
func (api *Client) GetGame(ctx context.Context, id int) (*GameDetailed, error) {
	path := fmt.Sprintf("/games/%d", id)

	var platform GameDetailed

	if err := api.get(ctx, path, nil, &platform); err != nil {
		return nil, err
	}

	return &platform, nil
}

// GetGameAchievements returns a list of game achievements
func (api *Client) GetGameAchievements(ctx context.Context, id int) ([]*Achievement, int, error) {
	path := fmt.Sprintf("/games/%d/achievements", id)

	var response struct {
		Results []*Achievement `json:"results"`
		Count   int            `json:"count"`
	}

	if err := api.get(ctx, path, nil, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameMovies returns a list of game trailers
func (api *Client) GetGameMovies(ctx context.Context, id int) ([]*Movie, int, error) {
	path := fmt.Sprintf("/games/%d/movies", id)

	var response struct {
		Results []*Movie `json:"results"`
		Count   int      `json:"count"`
	}

	if err := api.get(ctx, path, nil, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameReddit returns a list of most recent posts from the game's subreddit
func (api *Client) GetGameReddit(ctx context.Context, id int) ([]*Reddit, int, error) {
	path := fmt.Sprintf("/games/%d/reddit", id)

	var response struct {
		Results []*Reddit `json:"results"`
		Count   int       `json:"count"`
	}

	if err := api.get(ctx, path, nil, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameSuggested returns a list of visually similar games
func (api *Client) GetGameSuggested(ctx context.Context, id int) ([]*Game, int, error) {
	path := fmt.Sprintf("/games/%d/suggested", id)

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := api.get(ctx, path, nil, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameTwitch returns streams on Twitch associated with the game
func (api *Client) GetGameTwitch(ctx context.Context, id int) ([]*Twitch, int, error) {
	path := fmt.Sprintf("/games/%d/twitch", id)

	var response struct {
		Results []*Twitch `json:"results"`
		Count   int       `json:"count"`
	}

	if err := api.get(ctx, path, nil, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}

// GetGameYoutube returns videos from YouTube associated with the game
func (api *Client) GetGameYoutube(ctx context.Context, id int) ([]*Youtube, int, error) {
	path := fmt.Sprintf("/games/%d/youtube", id)

	var response struct {
		Results []*Youtube `json:"results"`
		Count   int        `json:"count"`
	}

	if err := api.get(ctx, path, nil, &response); err != nil {
		return nil, 0, err
	}

	return response.Results, response.Count, nil
}
