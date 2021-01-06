package rawg

func (suite *RAWGTestSuite) TestGetDevTeam() {
	team, total, err := suite.client.GetGameDevelopmentTeam(23, 1, 2, "-name")
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(team))
}

func (suite *RAWGTestSuite) TestGetDevTeamFailed() {
	suite.client.baseUrl = ""
	team, total, err := suite.client.GetGameDevelopmentTeam(23, 1, 2, "-name")
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(team))
}
