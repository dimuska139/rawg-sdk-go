package rawg

import "context"

func (suite *RAWGTestSuite) TestGetGameById() {
	game, err := suite.client.GetGame(context.Background(), "1")
	suite.NoError(err)
	suite.Equal("D/Generation HD", game.Name)
}

func (suite *RAWGTestSuite) TestGetGameBySlug() {
	game, err := suite.client.GetGame(context.Background(), "vampire-the-masquerade-bloodlines-2")
	suite.NoError(err)
	suite.Equal("Vampire: The Masquerade - Bloodlines 2", game.Name)
}

func (suite *RAWGTestSuite) TestGetGameFailed() {
	suite.client.baseUrl = ""
	game, err := suite.client.GetGame(context.Background(), "1")
	suite.Error(err)
	suite.Nil(game)
}
