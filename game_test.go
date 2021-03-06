package rawg

func (suite *RAWGTestSuite) TestGetGame() {
	game, err := suite.client.GetGame(1)
	suite.NoError(err)
	suite.Equal("D/Generation HD", game.Name)
}

func (suite *RAWGTestSuite) TestGetGameFailed() {
	suite.client.baseUrl = ""
	game, err := suite.client.GetGame(1)
	suite.Error(err)
	suite.Nil(game)
}
