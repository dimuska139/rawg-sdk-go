package rawg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetGames returns a list of games
func (api *Client) GetGames(filter *GamesFilter) ([]*Game, int, error) {
	path := "/games"
	body, err := api.newRequest(path, http.MethodGet, filter.GetParams())

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetGameAdditions returns a list of DLC's for the game, GOTY and other editions, companion apps, etc
func (api *Client) GetGameAdditions(gameID int, page int, pageSize int) ([]*Game, int, error) {
	path := fmt.Sprintf("/games/%d/additions", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetGameDevelopmentTeam returns a list of individual creators that were part of the development team
func (api *Client) GetGameDevelopmentTeam(gameID int, page int, pageSize int, ordering string) ([]*GameDeveloper, int, error) {
	path := fmt.Sprintf("/games/%d/development-team", gameID)
	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}

	if ordering != "" {
		data["ordering"] = ordering
	}

	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*GameDeveloper `json:"results"`
		Count   int              `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetGameSeries returns a list of games that are part of same series
func (api *Client) GetGameSeries(gameID int, page int, pageSize int) ([]*Game, int, error) {
	path := fmt.Sprintf("/games/%d/game-series", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetParentGames returns a list of parent games for DLC's and editions
func (api *Client) GetParentGames(gameID int, page int, pageSize int) ([]*Game, int, error) {
	path := fmt.Sprintf("/games/%d/parent-games", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetGameScreenshots returns a screenshots list for the game
func (api *Client) GetGameScreenshots(gameID int, page int, pageSize int) ([]*Screenshot, int, error) {
	path := fmt.Sprintf("/games/%d/screenshots", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*Screenshot `json:"results"`
		Count   int           `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetGameStores returns links to the stores that sell the game
func (api *Client) GetGameStores(gameID int, page int, pageSize int) ([]*GameStore, int, error) {
	path := fmt.Sprintf("/games/%d/stores", gameID)

	data := map[string]interface{}{
		"page":      fmt.Sprint(page),
		"page_size": fmt.Sprint(pageSize),
	}
	body, err := api.newRequest(path, http.MethodGet, data)

	if err != nil {
		return nil, 0, err
	}

	var response struct {
		Results []*GameStore `json:"results"`
		Count   int          `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, 0, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, response.Count, nil
}

// GetGame returns details of the game
func (api *Client) GetGame(id int) (*GameDetailed, error) {
	path := fmt.Sprintf("/games/%d", id)
	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var platform GameDetailed

	if err := json.Unmarshal(body, &platform); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &platform, nil
}

// GetGameAchievements returns a list of game achievements
func (api *Client) GetGameAchievements(id int) ([]*Achievement, error) {
	path := fmt.Sprintf("/games/%d/achievements", id)

	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var response struct {
		Results []*Achievement `json:"results"`
		Count   int            `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, nil
}

// GetGameMovies returns a list of game trailers
func (api *Client) GetGameMovies(id int) ([]*Movie, error) {
	path := fmt.Sprintf("/games/%d/movies", id)

	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var response struct {
		Results []*Movie `json:"results"`
		Count   int      `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, nil
}

// GetGameReddit returns a list of most recent posts from the game's subreddit
func (api *Client) GetGameReddit(id int) ([]*Reddit, error) {
	path := fmt.Sprintf("/games/%d/reddit", id)

	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var response struct {
		Results []*Reddit `json:"results"`
		Count   int       `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, nil
}

// GetGameSuggested returns a list of visually similar games
func (api *Client) GetGameSuggested(id int) ([]*Game, error) {
	path := fmt.Sprintf("/games/%d/suggested", id)

	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var response struct {
		Results []*Game `json:"results"`
		Count   int     `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, nil
}

// GetGameTwitch returns streams on Twitch associated with the game
func (api *Client) GetGameTwitch(id int) ([]*Twitch, error) {
	path := fmt.Sprintf("/games/%d/twitch", id)

	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var response struct {
		Results []*Twitch `json:"results"`
		Count   int       `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, nil
}

// GetGameYoutube returns videos from YouTube associated with the game
func (api *Client) GetGameYoutube(id int) ([]*Youtube, error) {
	path := fmt.Sprintf("/games/%d/youtube", id)

	body, err := api.newRequest(path, http.MethodGet, nil)

	if err != nil {
		return nil, err
	}

	var response struct {
		Results []*Youtube `json:"results"`
		Count   int        `json:"count"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &RawgError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return response.Results, nil
}
