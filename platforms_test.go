package rawg

import "context"

func (suite *RAWGTestSuite) TestGetPlatforms() {
	platforms, total, err := suite.client.GetPlatforms(context.Background(), 1, 2, "-name")
	suite.NoError(err)
	suite.NotEqual(0, len(platforms))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetPlatformsFailed() {
	suite.client.baseUrl = ""
	platforms, total, err := suite.client.GetPlatforms(context.Background(), 1, 2, "-name")
	suite.Error(err)
	suite.Equal(0, len(platforms))
	suite.Equal(0, total)
}
