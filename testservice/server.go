package testservice

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fasibio/autogql/testservice/graph"
	"github.com/fasibio/autogql/testservice/graph/db"
	"github.com/fasibio/autogql/testservice/graph/model"
	"gorm.io/gorm"
)

const defaultPort = "8432"

func StartServer(dbCon *gorm.DB) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dborm := db.NewAutoGqlDB(dbCon)
	dborm.Init()
	db.AddAddHook[model.Todo, model.TodoInput, model.AddTodoPayload](&dborm, "AddTodo", AddTodoHook{})
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{Sql: &dborm},
		Directives: graph.DirectiveRoot{
			VALIDATE: func(ctx context.Context, obj interface{}, next graphql.Resolver, value string) (res interface{}, err error) {
				return next(ctx)
			},
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
