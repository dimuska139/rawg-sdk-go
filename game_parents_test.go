package rawg

func (suite *RAWGTestSuite) TestGetParentGames() {
	games, total, err := suite.client.GetParentGames(34, 1, 2)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(games))
}

func (suite *RAWGTestSuite) TestGetParentGamesFailed() {
	suite.client.baseUrl = ""
	games, total, err := suite.client.GetParentGames(34, 1, 2)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(games))
}
