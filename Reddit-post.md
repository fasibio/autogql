# AutoGql:  GraphQL GORM CRUD generator library

Are you also tired of writing the same CRUD code over and over again? Yepp, same here - so i worked the last 6 months on a gqlgen plugin to automate all the boring parts!

Meet [AutoGql](https://github.com/fasibio/autogql)!

I also tried the entgql plugin, but wanted to stay schema first, so i added just a few new compile time directives to add and control the resolver generation for your entities.

Define your Types with all Gorm features inside GraphQL schema definition: 

```
type Group @SQL{
  id: ID! @SQL_PRIMARY @SQL_GORM(value: "autoIncrement")
  name: String!
  users: [User!]! @SQL_GORM(value:"many2many:group_users")
  shoppingLists: [ShoppingList]
  creator: User! @SQL_SKIP_MUTATION
  creatorID: ID! @SQL_SKIP_MUTATION
  createdAt: Time
  updatedAt: Time
  deletedAt: Time
}

type User @SQL{
  id: ID! @SQL_PRIMARY @SQL_GORM(value: "autoIncrement")
  nick_name: String!
  groups: [Group!] @SQL_GORM(value:"many2many:group_users")
  createdAt: Time
  updatedAt: Time
  deletedAt: Time
}

```
and use all Database which supported by Gorm (Postgres, SQLite, ...). 

You decide which Types are under "control" of [AutoGql](https://github.com/fasibio/autogql) and which not. 