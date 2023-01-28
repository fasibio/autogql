# AutoGql

## About
GraphQL Gorm CRUD Generator. 

Its a plugin for [99designs/gqlgen](https://github.com/99designs/gqlgen).

It helps you to make the CRUD-functionalities fast and let you focus to the real spezial thinks of your business.

## How to setup

- Follow the steps from [Gqlgen](gqlgen.com)
- Create a folder ```plugin``` and add a main.go inside. 
- Copy Content: 
```golang
package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/fasibio/autogql"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()

	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	sqlPlugin, muateHookPlugin := autogql.NewAutoGqlPlugin()
	err = api.Generate(cfg, api.AddPlugin(sqlPlugin), api.ReplacePlugin(muateHookPlugin))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
```
- install dependencies
```bash
go mod tidy
```
- add a example autogql struct to ```schema.graphqls```

```gql
type Company @SQL{
  id: Int! @SQL_PRIMARY
  Name: String!
}
```

- now you have to create gqlgen generation **always** with: 

```bash
go run plugin/main.go
```

- add sql entity to Resolver struct at ```resolver.go```
```go
type Resolver struct {
	Sql *db.AutoGqlDB // this is the new line the package db is autogenerate by this plugin
}
```

- add SQL Gorm Connection(here used SQLite as Example) to ```server.go```:

```go

import (
  // ... more imports
  "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
func main() {
  dbCon, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
		panic(err)
	}
  dborm := db.NewAutoGqlDB(dbCon)
  dborm.Init()

  srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Sql: &dborm} //.... <- here set dborm to resolver
}

```

# Directives

Each type you will be managed by this Lib add  ```@SQL```
```gql
type Company @SQL{
```

It will autogenerate Queries and Mutations. Also it will create resolvers and fill with gorm Database code. 

Description: 

```gql

	input SqlCreateExtension {
		value: Boolean! # active this query or mutation
		directiveExt: [String!] # add directive to query or mutation
	}

	input SqlMutationParams {
		add: SqlCreateExtension
		update: SqlCreateExtension
		delete: SqlCreateExtension
		directiveExt: [String!] # add directive to all mutation
	}

	input SqlQueryParams {
		get: SqlCreateExtension
		query: SqlCreateExtension
		directiveExt: [String!] # dd directive to all mutations
	}
	directive @SQL(query:SqlQueryParams, mutation: SqlMutationParams ) on OBJECT
	directive @SQL_PRIMARY on FIELD_DEFINITION
	directive @SQL_INDEX on FIELD_DEFINITION

	directive @SQL_GORM (value: String)on FIELD_DEFINITION # each gorm command ==> not all useable at the moment pls open issue if you find one

```



# Hooks 

Each Query and Mutation can be manipulated over Hooks. 

All Hooksdescription are written at db/db_gen.go

See [autogql_example](https://github.com/fasibio/autogql_example) 
 - [hooks.go](https://github.com/fasibio/autogql_example/blob/main/hooks.go)
 - And to include : [server.go](https://github.com/fasibio/autogql_example/blob/main/server.go#L29-L33)





