package rawg

func (suite *RAWGTestSuite) TestGetPublisher() {
	publisher, err := suite.client.GetPublisher(3)
	suite.NoError(err)
	suite.Equal("Juicy Beast Studio", publisher.Name)
}

func (suite *RAWGTestSuite) TestGetPublisherFailed() {
	suite.client.baseUrl = ""
	publisher, err := suite.client.GetPublisher(3)
	suite.Error(err)
	suite.Nil(publisher)
}
