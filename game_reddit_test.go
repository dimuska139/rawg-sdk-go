package rawg

func (suite *RAWGTestSuite) TestGetRedditTopics() {
	topics, total, err := suite.client.GetGameReddit(25)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(topics))
}

func (suite *RAWGTestSuite) TestGetRedditTopicsFailed() {
	suite.client.baseUrl = ""
	topics, total, err := suite.client.GetGameReddit(25)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(topics))
}
