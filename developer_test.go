package rawg

func (suite *RAWGTestSuite) TestGetDeveloper() {
	developer, err := suite.client.GetDeveloper(1)
	suite.NoError(err)
	suite.Equal("D3 Publisher of America", developer.Name)
}

func (suite *RAWGTestSuite) TestGetDeveloperFailed() {
	suite.client.baseUrl = ""
	developer, err := suite.client.GetDeveloper(1)
	suite.Error(err)
	suite.Nil(developer)
}
