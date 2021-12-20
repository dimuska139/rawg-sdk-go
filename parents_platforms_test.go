package rawg

import "context"

func (suite *RAWGTestSuite) TestGetParentsPlatforms() {
	parentsPlatforms, total, err := suite.client.GetParentsPlatforms(context.Background(), 1, 2, "-name")
	suite.NoError(err)
	suite.NotEqual(0, len(parentsPlatforms))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetParentsPlatformsFailed() {
	suite.client.baseUrl = ""
	parentsPlatforms, total, err := suite.client.GetParentsPlatforms(context.Background(), 1, 2, "-name")
	suite.Error(err)
	suite.Equal(0, len(parentsPlatforms))
	suite.Equal(0, total)
}
