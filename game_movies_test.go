package rawg

func (suite *RAWGTestSuite) TestGetMovies() {
	movies, total, err := suite.client.GetGameMovies(23)
	suite.NoError(err)
	suite.NotEqual(0, total)
	suite.NotEqual(0, len(movies))
}

func (suite *RAWGTestSuite) TestGetMoviesFailed() {
	suite.client.baseUrl = ""
	movies, total, err := suite.client.GetGameMovies(23)
	suite.Error(err)
	suite.Equal(0, total)
	suite.Equal(0, len(movies))
}
