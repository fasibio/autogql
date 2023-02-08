# AutoGql

AutoGQL is a GraphQL GORM CRUD generator. It's a plugin for 99designs/gqlgen that helps you quickly create CRUD functionality for your application so you can focus on the unique aspects of your business

## Setup

1. Follow the steps from [Gqlgen](gqlgen.com)
2. Create a folder ```plugin``` and add a ```main.go``` inside. 
3. Copy the following code into ```main.go```: 
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
4. Install dependencies
```bash
go mod tidy
```
5. Add a example autogql struct to ```schema.graphqls```

```gql
type Company @SQL{
  id: Int! @SQL_PRIMARY
  Name: String!
}
```

6. Run the following command to generate the GQLgen code:

```bash
go run plugin/main.go
```

7. Add SQL entity to ```Resolver``` struct at ```resolver.go```
```go
type Resolver struct {
	Sql *db.AutoGqlDB // this is the new line the package db is autogenerate by this plugin
}
```

8. Add SQL GORM Connection(SQLite used as an example) to ```server.go```:

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
Add the ```@SQL``` directive to each type that you want managed by AutoGQL.

```graphql
...
type Company @SQL{
	id: Int! @SQL_PRIMARY
	...
}
...
```
It will autogenerate Queries and Mutations based on your GraphQL schema. Also, it will create resolvers and fill them with GORM Database code.

Description: 

```graphql

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
		directiveExt: [String!] # add directive to all mutations
	}
	directive @SQL(query:SqlQueryParams, mutation: SqlMutationParams ) on OBJECT
	directive @SQL_PRIMARY on FIELD_DEFINITION
	directive @SQL_INDEX on FIELD_DEFINITION

	directive @SQL_GORM (value: String)on FIELD_DEFINITION # each gorm command ==> not all useable at the moment pls open issue if you find one

	scalar Time #activated for createdAt, deletedAt, updatedAt etc

```

If an field has Tag autoIncrement it will be not include into patch and input Types. 
f.E.: 
```graphql
type Cat @SQL{
  id: Int! @SQL_PRIMARY @SQL_GORM(value: "autoIncrement") // id will be removed from add and patch types
  name: String!
  age: Int
  userID: Int!
  alive: Boolean @SQL_GORM(value: "default:true")
}

```


# Hooks 

You can manipulate each Query and Mutation through Hooks. The Hooks descriptions are written in db/db_gen.go. For more information, see the [autogql_example](https://github.com/fasibio/autogql_example) repository:


 - [hooks.go](https://github.com/fasibio/autogql_example/blob/main/hooks.go)
 - And to include : [server.go](https://github.com/fasibio/autogql_example/blob/main/server.go#L29-L33)





