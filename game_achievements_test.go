package rawg

import "context"

func (suite *RAWGTestSuite) TestGetAchievements() {
	achievements, total, err := suite.client.GetGameAchievements(context.Background(), 23)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(achievements))
}

func (suite *RAWGTestSuite) TestGetAchievementsFailed() {
	suite.client.baseUrl = ""
	achievements, total, err := suite.client.GetGameAchievements(context.Background(), 23)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(achievements))
}
