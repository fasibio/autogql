package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/Khan/genqlient/graphql"
	"github.com/fasibio/autogql/testservice"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QueryTestSuite struct {
	suite.Suite
	Client graphql.Client
}

func (suite *QueryTestSuite) SetupTest() {
	go testservice.StartServer()
	h := http.Client{}
	suite.Client = graphql.NewClient("http://localhost:8432/query", &h)

}

func (suite *QueryTestSuite) TestIntrospection() {
	resp, err := IntrospectionQuery(context.TODO(), suite.Client)
	assert.Nil(suite.T(), err)
	snaps.MatchSnapshot(suite.T(), resp)
}

func TestQueryTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(QueryTestSuite))
}
