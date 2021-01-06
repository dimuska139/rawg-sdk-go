package rawg

func (suite *RAWGTestSuite) TestGetTwitchVideos() {
	videos, total, err := suite.client.GetGameTwitch(1)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(videos))
}

func (suite *RAWGTestSuite) TestGetTwitchFailed() {
	suite.client.baseUrl = ""
	videos, total, err := suite.client.GetGameTwitch(1)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(videos))
}
