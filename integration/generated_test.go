package integration

import (
	"context"
	"net/http"
	"testing"
	"time"

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

func (suite *QueryTestSuite) SetupSuite() {
	go testservice.StartServer()
	time.Sleep(15 + time.Second)
}

func (suite *QueryTestSuite) SetupTest() {
	h := http.Client{}
	suite.Client = graphql.NewClient("http://localhost:8432/query", &h)

}

func (suite *QueryTestSuite) TestIntrospection() {
	resp, err := IntrospectionQuery(context.TODO(), suite.Client)
	assert.Nil(suite.T(), err)
	snaps.MatchSnapshot(suite.T(), resp)
}

func (suite *QueryTestSuite) AddCompanies(t *testing.T) {
	resp, err := addCompanies(context.TODO(), suite.Client, []*CompanyInput{
		{
			Name: "TestCompany1",
		},
		{
			Name: "TestCompany2",
		},
	})
	assert.Nil(t, err)
	snaps.MatchSnapshot(t, resp)
}

func (suite *QueryTestSuite) AddUsers(t *testing.T) {
	resp, err := addUsers(context.TODO(), suite.Client, []*UserInput{
		{
			Name:      "Jan",
			CompanyID: getPointerOf(1),
		},
		{
			Name:      "Klaas",
			CompanyID: getPointerOf(1),
		},
		{
			Name:      "Peter",
			CompanyID: getPointerOf(1),
		},
		{
			Name:      "Schmadel",
			CompanyID: getPointerOf(1),
		},
		{
			Name:      "Jan",
			CompanyID: getPointerOf(2),
		},
		{
			Name:      "Boris",
			CompanyID: getPointerOf(2),
		},
	})
	assert.Nil(t, err)
	snaps.MatchSnapshot(t, resp)
}

func (suite *QueryTestSuite) AddCats(t *testing.T) {
	resp, err := addCats(context.TODO(), suite.Client, []*CatInput{
		{
			Name:     "Pussy",
			BirthDay: time.Now(),
			UserID:   1,
			Alive:    getPointerOf(true),
		},
		{
			Name:     "Schnuffel",
			BirthDay: time.Now(),
			UserID:   2,
		},
		{
			Name:     "Mauz",
			BirthDay: time.Now(),
			UserID:   3,
			Alive:    getPointerOf(false),
		},
		{
			Name:     "Mi jau",
			BirthDay: time.Now(),
			UserID:   4,
			Alive:    getPointerOf(true),
		},
	})
	assert.Nil(t, err)
	snaps.MatchSnapshot(t, resp)
}

func (suite *QueryTestSuite) AddTodos(t *testing.T) {
	resp, err := addTodos(context.TODO(), suite.Client, []*TodoInput{
		{
			Name: "Task 1",
		},
		{
			Name: "Task 2",
		},
		{
			Name: "Task 3",
		},
		{
			Name: "Task 4",
		},
		{
			Name: "Task 5",
		},
	})
	assert.Nil(t, err)
	snaps.MatchSnapshot(t, resp)
}

func getPointerOf[T any](v T) *T {
	return &v
}

func (suite *QueryTestSuite) AddUsers2Todo(t *testing.T) {

	resp, err := addUser2Todo(context.TODO(), suite.Client, &UserRef2TodosInput{
		Set: []string{"1", "3"},
		Filter: &TodoFiltersInput{
			Or: []*TodoFiltersInput{
				{Id: &IDFilterInput{Eq: getPointerOf("1")}},
				{Id: &IDFilterInput{Eq: getPointerOf("3")}},
				{Id: &IDFilterInput{Eq: getPointerOf("5")}},
			},
		},
	})
	assert.Nil(t, err)
	snaps.MatchSnapshot(t, resp)
}

type queryTesterFunction = func() (any, error)

func queryTester(f queryTesterFunction) func(*testing.T) {
	return func(t *testing.T) {
		a, err := f()
		assert.Nil(t, err)
		snaps.MatchSnapshot(t, a)
	}

}

func (suite *QueryTestSuite) TestComplexCombination() {
	suite.T().Run("addCompanies", suite.AddCompanies)
	suite.T().Run("addUsers", suite.AddUsers)
	suite.T().Run("addCats", suite.AddCats)
	suite.T().Run("addTodos", suite.AddTodos)
	suite.T().Run("addUsers2Todo", suite.AddUsers2Todo)
	suite.T().Run("allUserFromCompany => TestCompany1, offset 0", queryTester(func() (any, error) {
		offset := 0
		return allUserFromCompany(context.Background(), suite.Client, "TestCompany1", &offset)
	}))
}

func TestQueryTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(QueryTestSuite))
}
