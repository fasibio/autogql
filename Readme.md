# AutoGql

AutoGQL is a GraphQL GORM CRUD generator. It's a plugin for 99designs/gqlgen that helps you quickly create CRUD functionality for your application so you can focus on the unique aspects of your business

## Setup

1. Follow the steps from [Gqlgen](https://gqlgen.com)
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
	sqlPlugin, muateHookPlugin := autogql.NewAutoGqlPlugin(cfg)
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
# How to Use Autogql
Autogql is designed to work similarly to [gorm](https://gorm.io/) for declaring database tables and relations. To get started, it's helpful to familiarize yourself with how [gorm models](https://gorm.io/docs/models.html) work:
## Little introduction
At gorm you describe Database tables and relations like this
```go
type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name string
}

func (u User) Calculation() int {
	return 10
}
```
With Autogql, you can describe GraphQL schemas and directives in the same way:

```graphql

type User @SQL{
	id: Int! @SQL_PRIMARY @SQL_GORM(value: "autoIncrement")
	createdAt: Time
	updatedAt: Time
	deletedAt: SoftDelete #activate softdelete
	name: String
	calculation: Int! @SQL_GORM(value: "-")
}

```
In this example, the `@SQL` directive tells Autogql that this type should be mapped to a GORM model. The `@SQL_GORM` directive specifies the corresponding GORM tag for the field.

You can find more examples of Autogql schema descriptions in the [test schema](./testservice/graph/schema.graphqls) file. Autogql also allows you to define relationships between models using GORM's foreign key syntax:


```graphql

type Cat @SQL(order: 4){
  id: ID! @SQL_PRIMARY @SQL_GORM(value: "autoIncrement")
  name: String!
  birthDay: Time!
  age: Int @SQL_GORM(value:"-")
  userID: Int! #<--- foreign Key to user
  alive: Boolean @SQL_GORM(value: "default:true")
}

type User @SQL(order: 2){
  id: ID! @SQL_PRIMARY @SQL_GORM(value: "autoIncrement")
  name: String!
  createdAt: Time
  updatedAt: Time
  deletedAt: SoftDelete #activate softdelete
  cat: Cat @SQL_GORM(value:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;")# <--- cat defintion. SQL_GORM not needed for relation only a constraint line
  companyID: Int
  company: Company
  smartPhones: [SmartPhone]
}

```
In this example, the `userID` field in the `Cat` type is a foreign key that points to the `id` field in the `User` type. The `@SQL_GORM` directive can be used to specify GORM tags for relationships, as well as other constraints.

To work with queries and mutations, Autogql will automatically generate them for you, as well as create resolvers and fill in GORM database code. Additionally, you can manipulate each query and mutation over hooks, as described in the `db/db_gen.go` file. For an example of how to include hooks, check out the [autogql_example](https://github.com/fasibio/autogql_example)  repository and the [hooks.go](https://github.com/fasibio/autogql_example/blob/main/hooks.go) file.


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
	# order to define relations (if Type A(order: 1) have a releation to type B(order: 2)
	directive @SQL(order: Int, query:SqlQueryParams, mutation: SqlMutationParams ) on OBJECT 

	# database primary key
	directive @SQL_PRIMARY on FIELD_DEFINITION

	# database index
	directive @SQL_INDEX on FIELD_DEFINITION

	# each gorm command ==> not all useable at the moment pls open issue if you find one
	directive @SQL_GORM (value: String)on FIELD_DEFINITION 

	# to remove this value from input and patch generated Inputs
	directive @SQL_SKIP_MUTATION on FIELD_DEFINITION 

	# to add a tag to input go struct
	directive @SQL_INPUTTYPE_TAGS (value: [String!]) on FIELD_DEFINITION 

	#to add a directive to input graphql type (directive have to be decelerated with INPUT_FIELD_DEFINITION or INPUT_OBJECT )
	directive @SQL_INPUTTYPE_DIRECTIVE (value: [String!]) on FIELD_DEFINITION | OBJECT


	scalar Time #activated for createdAt, updatedAt etc
	scalar SoftDelete # for deleteAt field for softdelete

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


# Scalar ID as autoIncrement Primary key
To use ID as autoIncrement Primary key you have to update ```gqlgen.yml```
from : 
```yaml
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
```
to 
```yaml
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.IntID
```

to use Int instand of String


# Hooks 

You can manipulate each Query and Mutation through Hooks. The Hooks descriptions are written in db/db_gen.go. For more information, see the [autogql_example](https://github.com/fasibio/autogql_example) repository:


 - [hooks.go](https://github.com/fasibio/autogql_example/blob/main/hooks.go)
 - And to include : [server.go](https://github.com/fasibio/autogql_example/blob/main/server.go#L29-L33)

A default hook implementation will also be added to db package so you did not have to implement all hook functions: 

```go
type CompanyGetHook struct {
	db.DefaultGetHook[model.Company, int]
}

func (g CompanyGetHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	// your implementation
	return db, nil 
}
```



