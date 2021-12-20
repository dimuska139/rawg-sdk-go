package rawg

import "context"

func (suite *RAWGTestSuite) TestGetCreatorRoles() {
	roles, total, err := suite.client.GetCreatorRoles(context.Background(), 1, 2)
	suite.NoError(err)
	suite.Equal(Role{
		ID:   1,
		Name: "writer",
		Slug: "writer",
	}, *roles[0])
	suite.NotEqual(0, total)
}

func (suite *RAWGTestSuite) TestGetCreatorRolesFailed() {
	suite.client.baseUrl = ""
	_, _, err := suite.client.GetCreatorRoles(context.Background(), 1, 2)
	suite.Error(err)
}
