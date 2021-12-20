package rawg

import "context"

func (suite *RAWGTestSuite) TestGetPublishers() {
	publishers, total, err := suite.client.GetPublishers(context.Background(), 1, 2)
	suite.NoError(err)
	suite.NotEqual(0, len(publishers))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetPublishersFailed() {
	suite.client.baseUrl = ""
	publishers, total, err := suite.client.GetPublishers(context.Background(), 1, 2)
	suite.Error(err)
	suite.Equal(0, len(publishers))
	suite.Equal(0, total)
}
