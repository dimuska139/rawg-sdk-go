package rawg

import "context"

func (suite *RAWGTestSuite) TestGetGames() {
	filter := NewGamesFilter().
		SetPage(1).
		SetPageSize(2).
		SetSearch("gta5")
	games, total, err := suite.client.GetGames(context.Background(), filter)
	suite.NoError(err)
	suite.Equal(2, len(games))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetGamesFailed() {
	suite.client.baseUrl = ""
	filter := NewGamesFilter().
		SetPage(1).
		SetPageSize(2).
		SetSearch("gta5")
	games, total, err := suite.client.GetGames(context.Background(), filter)
	suite.Error(err)
	suite.Equal(0, len(games))
	suite.Equal(0, total)
}
