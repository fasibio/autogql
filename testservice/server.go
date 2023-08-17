package testservice

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fasibio/autogql/testservice/graph"
	"github.com/fasibio/autogql/testservice/graph/db"
	"github.com/fasibio/autogql/testservice/graph/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

const defaultPort = "8432"

var validate *validator.Validate

func StartServer(dbCon *gorm.DB) {
	validate = validator.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dborm := db.NewAutoGqlDB(dbCon)
	err := dborm.Init()
	if err != nil {
		panic(err)
	}
	db.AddAddHook[model.Todo, model.TodoInput, model.AddTodoPayload](&dborm, db.AddTodo, AddTodoHook{})
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{Sql: &dborm},
		Directives: graph.DirectiveRoot{
			VALIDATE: ValidateDirective,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
