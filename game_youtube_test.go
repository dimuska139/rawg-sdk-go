package rawg

import "context"

func (suite *RAWGTestSuite) TestGetYoutubeVideos() {
	videos, total, err := suite.client.GetGameYoutube(context.Background(), 1)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(videos))
}

func (suite *RAWGTestSuite) TestGetYoutubeVideosFailed() {
	suite.client.baseUrl = ""
	videos, total, err := suite.client.GetGameYoutube(context.Background(), 1)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(videos))
}
