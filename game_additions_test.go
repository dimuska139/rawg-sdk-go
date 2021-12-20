package rawg

import "context"

func (suite *RAWGTestSuite) TestGetGameAdditions() {
	additions, total, err := suite.client.GetGameAdditions(context.Background(), 123, 1, 2)
	suite.NoError(err)
	suite.Equal(2, len(additions))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetGameAdditionsFailed() {
	suite.client.baseUrl = ""
	additions, total, err := suite.client.GetGameAdditions(context.Background(), 123, 1, 2)
	suite.Error(err)
	suite.Equal(0, len(additions))
	suite.Equal(0, total)
}
