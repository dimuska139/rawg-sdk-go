package rawg

import "context"

func (suite *RAWGTestSuite) TestGetStores() {
	stores, total, err := suite.client.GetStores(context.Background(), 1, 2, "-name")
	suite.NoError(err)
	suite.NotEqual(0, len(stores))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetStoresFailed() {
	suite.client.baseUrl = ""
	stores, total, err := suite.client.GetStores(context.Background(), 1, 2, "-name")
	suite.Error(err)
	suite.Equal(0, len(stores))
	suite.Equal(0, total)
}
