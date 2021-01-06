package rawg

func (suite *RAWGTestSuite) TestGetCreator() {
	creator, err := suite.client.GetCreator(1)
	suite.NoError(err)
	suite.Equal("Michael Unsworth", creator.Name)
}

func (suite *RAWGTestSuite) TestGetCreatorFailed() {
	suite.client.baseUrl = ""
	creator, err := suite.client.GetCreator(1)
	suite.Error(err)
	suite.Nil(creator)
}
