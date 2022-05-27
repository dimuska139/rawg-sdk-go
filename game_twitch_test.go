package rawg

import "context"

func (suite *RAWGTestSuite) TestGetTwitchVideos() {
	videos, total, err := suite.client.GetGameTwitch(context.Background(), 10213)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(videos))
}

func (suite *RAWGTestSuite) TestGetTwitchFailed() {
	suite.client.baseUrl = ""
	videos, total, err := suite.client.GetGameTwitch(context.Background(), 6256)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(videos))
}
