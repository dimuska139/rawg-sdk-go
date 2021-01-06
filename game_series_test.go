package rawg

func (suite *RAWGTestSuite) TestGetSeries() {
	series, total, err := suite.client.GetGameSeries(21, 1, 2)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(series))
}

func (suite *RAWGTestSuite) TestGetSeriesFailed() {
	suite.client.baseUrl = ""
	series, total, err := suite.client.GetGameSeries(21, 1, 2)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(series))
}
