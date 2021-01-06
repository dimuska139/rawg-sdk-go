package rawg

func (suite *RAWGTestSuite) TestGetCreators() {
	creators, total, err := suite.client.GetCreators(1, 2)
	suite.NoError(err)
	suite.Equal(2, len(creators))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetCreatorsFailed() {
	suite.client.baseUrl = ""
	creators, total, err := suite.client.GetCreators(1, 2)
	suite.Error(err)
	suite.Equal(0, len(creators))
	suite.Equal(0, total)
}
