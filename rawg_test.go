package rawg

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

type RAWGTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *RAWGTestSuite) SetupSuite() {
	config := Config{
		ApiKey:   os.Getenv("RAWG_API_KEY"),
		Language: "en",
		Rps:      5,
	}
	suite.client = NewClient(http.DefaultClient, &config)
}

func (suite *RAWGTestSuite) SetupTest() {
	suite.client.baseUrl = "https://api.rawg.io/api"
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RAWGTestSuite))
}
