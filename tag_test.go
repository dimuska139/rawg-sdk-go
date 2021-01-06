package rawg

func (suite *RAWGTestSuite) TestGetTag() {
	tag, err := suite.client.GetTag(1)
	suite.NoError(err)
	suite.Equal("Survival", tag.Name)
}

func (suite *RAWGTestSuite) TestGetTagFailed() {
	suite.client.baseUrl = ""
	tag, err := suite.client.GetTag(1)
	suite.Error(err)
	suite.Nil(tag)
}
