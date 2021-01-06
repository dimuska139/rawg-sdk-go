package rawg

func (suite *RAWGTestSuite) TestGetTags() {
	stores, total, err := suite.client.GetTags(1, 2)
	suite.NoError(err)
	suite.NotEqual(0, len(stores))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetTagsFailed() {
	suite.client.baseUrl = ""
	stores, total, err := suite.client.GetTags(1, 2)
	suite.Error(err)
	suite.Equal(0, len(stores))
	suite.Equal(0, total)
}
