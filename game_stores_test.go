package rawg

func (suite *RAWGTestSuite) TestGetGameStores() {
	stores, total, err := suite.client.GetGameStores(23, 1, 2)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(stores))
}

func (suite *RAWGTestSuite) TestGetGameStoresFailed() {
	suite.client.baseUrl = ""
	stores, total, err := suite.client.GetGameStores(23, 1, 2)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(stores))
}
