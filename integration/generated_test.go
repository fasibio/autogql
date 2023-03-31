package integration

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/fasibio/autogql/testservice"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultServerPort = "8432"

type QueryTestSuite struct {
	suite.Suite
	Client graphql.Client
}

type DatabaseType = string

const (
	Sqlite   DatabaseType = "sqlite"
	MySql    DatabaseType = "mysql"
	Postgres DatabaseType = "postgres"
)

func getDatabaseTestSystem() (*gorm.DB, error) {
	t := os.Getenv("DATABASE_TYPE")
	conStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if conStr == "" {
		panic("environment DATABASE_CONNECTION_STRING missing")
	}
	if t == "" || t == Sqlite {
		return gorm.Open(sqlite.Open(conStr), &gorm.Config{})
	}
	switch t {
	case MySql:
		return gorm.Open(mysql.Open(conStr), &gorm.Config{})
	case Postgres:
		return gorm.Open(postgres.Open(conStr), &gorm.Config{})
	default:
		panic(fmt.Sprintf("%s not a valid Databasetype", t))
	}
}

func (suite *QueryTestSuite) SetupSuite() {
	dbCon, err := getDatabaseTestSystem()
	if err != nil {
		panic(err)
	}
	dbCon = dbCon.Debug()
	go testservice.StartServer(dbCon)
	time.Sleep(15 + time.Second)
}

func (suite *QueryTestSuite) SetupTest() {
	h := http.Client{}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultServerPort
	}
	suite.Client = graphql.NewClient(fmt.Sprintf("http://localhost:%s/query", port), &h)

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
	suite.T().Run("allUserFromCompany => TestCompany1, offset 2", queryTester(func() (any, error) {
		offset := 2
		return allUserFromCompany(context.Background(), suite.Client, "TestCompany1", &offset)
	}))
	suite.T().Run("allUserFromCompany => TestCompany2, offset 0", queryTester(func() (any, error) {
		offset := 0
		return allUserFromCompany(context.Background(), suite.Client, "TestCompany2", &offset)
	}))
	suite.T().Run("allUserWithACat", queryTester(func() (any, error) {
		return allUserWithACat(context.Background(), suite.Client)
	}))
	suite.T().Run("allTodosPartOfCompany => TestCompany1", queryTester(func() (any, error) {
		companyName := "TestCompany1"
		return allTodosPartOfCompany(context.Background(), suite.Client, &companyName)
	}))
	suite.T().Run("allTodosPartOfCompany => TestCompany2", queryTester(func() (any, error) {
		companyName := "TestCompany2"
		return allTodosPartOfCompany(context.Background(), suite.Client, &companyName)
	}))
	suite.T().Run("getUserById => 2", queryTester(func() (any, error) {
		return getUserById(context.Background(), suite.Client, "2")
	}))
	suite.T().Run("getUserById => 5", queryTester(func() (any, error) {
		return getUserById(context.Background(), suite.Client, "5")
	}))
	suite.T().Run("updateUserChangeCompany => user with id 2 from Company 1 to 2", queryTester(func() (any, error) {
		return updateUserChangeCompany(context.Background(), suite.Client, "2", 2)
	}))
	suite.T().Run("updateUserChangeCompanyByCatName => user with id 2 from Company 2 to 1 by cat name Schnuffel", queryTester(func() (any, error) {
		return updateUserChangeCompanyByCatName(context.Background(), suite.Client, "Schnuffel", 1)
	}))
	suite.T().Run("changeAllCatsToSameOwner => user id 6 and cat not called Schnuffel", queryTester(func() (any, error) {
		return changeAllCatsToSameOwnerButNotOneByName(context.Background(), suite.Client, 6, "Schnuffel")
	}))
	suite.T().Run("deleteUser => user with id 1", queryTester(func() (any, error) {
		return deleteUser(context.Background(), suite.Client, "1")
	}))
	suite.T().Run("deleteUser => user with id 5", queryTester(func() (any, error) {
		return deleteUser(context.Background(), suite.Client, "5")
	}))
	suite.T().Run("deleteUserByCatName => user 2 with cat Schnuffel", queryTester(func() (any, error) {
		return deleteUserByCatName(context.Background(), suite.Client, "Schnuffel")
	}))
	suite.T().Run("deleteUserByUserName => user 4 with name Schmadel", queryTester(func() (any, error) {
		return deleteUserByUserName(context.Background(), suite.Client, "Schmadel")
	}))
}

func TestQueryTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(QueryTestSuite))
}
