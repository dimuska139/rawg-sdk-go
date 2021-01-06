package rawg

func (suite *RAWGTestSuite) TestGetGenres() {
	genres, total, err := suite.client.GetGenres(1, 2, "-name")
	suite.NoError(err)
	suite.Equal(2, len(genres))
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetGenresFailed() {
	suite.client.baseUrl = ""
	genres, total, err := suite.client.GetGenres(1, 2, "-name")
	suite.Error(err)
	suite.Equal(0, len(genres))
	suite.Equal(0, total)
}
