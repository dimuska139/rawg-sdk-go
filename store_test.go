package rawg

func (suite *RAWGTestSuite) TestGetStore() {
	store, err := suite.client.GetStore(1)
	suite.NoError(err)
	suite.Equal("Steam", store.Name)
}

func (suite *RAWGTestSuite) TestGetStoreFailed() {
	suite.client.baseUrl = ""
	store, err := suite.client.GetStore(1)
	suite.Error(err)
	suite.Nil(store)
}
