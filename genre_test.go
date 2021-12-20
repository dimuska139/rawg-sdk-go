package rawg

import "context"

func (suite *RAWGTestSuite) TestGetGenre() {
	genre, err := suite.client.GetGenre(context.Background(), 1)
	suite.NoError(err)
	suite.Equal("Racing", genre.Name)
}

func (suite *RAWGTestSuite) TestGetGenreFailed() {
	suite.client.baseUrl = ""
	genre, err := suite.client.GetGenre(context.Background(), 1)
	suite.Error(err)
	suite.Nil(genre)
}
