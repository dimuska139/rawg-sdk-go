package rawg

func (suite *RAWGTestSuite) TestGetDevelopers() {
	developers, total, err := suite.client.GetDevelopers(1, 2)
	suite.NoError(err)
	suite.Equal(2, len(developers))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetDevelopersFailed() {
	suite.client.baseUrl = ""
	developers, total, err := suite.client.GetDevelopers(1, 2)
	suite.Error(err)
	suite.Equal(0, len(developers))
	suite.Equal(0, total)
}
