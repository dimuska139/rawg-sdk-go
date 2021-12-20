package rawg

import "context"

func (suite *RAWGTestSuite) TestGetScreenshots() {
	screenshots, total, err := suite.client.GetGameScreenshots(context.Background(), 23, 1, 2)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(screenshots))
}

func (suite *RAWGTestSuite) TestGetScreenshotsFailed() {
	suite.client.baseUrl = ""
	screenshots, total, err := suite.client.GetGameScreenshots(context.Background(), 23, 1, 2)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(screenshots))
}
