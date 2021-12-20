package rawg

import "context"

func (suite *RAWGTestSuite) TestGetDeveloper() {
	developer, err := suite.client.GetDeveloper(context.Background(), 1612)
	suite.NoError(err)
	suite.Equal("Valve Software", developer.Name)
}

func (suite *RAWGTestSuite) TestGetDeveloperFailed() {
	suite.client.baseUrl = ""
	developer, err := suite.client.GetDeveloper(context.Background(), 1612)
	suite.Error(err)
	suite.Nil(developer)
}
