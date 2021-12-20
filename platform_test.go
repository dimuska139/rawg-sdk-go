package rawg

import "context"

func (suite *RAWGTestSuite) TestGetPlatform() {
	platform, err := suite.client.GetPlatform(context.Background(), 1)
	suite.NoError(err)
	suite.Equal("Xbox One", platform.Name)
}

func (suite *RAWGTestSuite) TestGetPlatformFailed() {
	suite.client.baseUrl = ""
	platform, err := suite.client.GetPlatform(context.Background(), 1)
	suite.Error(err)
	suite.Nil(platform)
}
