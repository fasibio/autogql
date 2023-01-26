{{- $root := .}}


input IDFilterInput {
  and: [ID]
  or: [ID]
  not: IDFilterInput
  eq: ID
  ne: ID
  null: Boolean
  notNull: Boolean
  in: [ID]
  notin: [ID]
}

input StringFilterInput {
  and: [String]
  or: [String]
  not: StringFilterInput
  eq: String
  eqi: String 
  ne: String
  startsWith: String
  endsWith: String
  contains: String
  notContains: String
  containsi: String
  notContainsi: String
  null: Boolean
  notNull: Boolean
  in: [String]
  notIn: [String]
}

input IntFilterInput {
  and: [Int]
  or: [Int]
  not: IntFilterInput
  eq: Int
  ne: Int
  gt: Int
  gte: Int
  lt: Int
  lte: Int
  null: Boolean
  notNull: Boolean
  in: [Int]
  notIn: [Int]
  between: IntFilterBetween
}

input IntFilterBetween{
  start: Int!
  end: Int!
}

input BooleanFilterInput{
  and: [Boolean]
  or: [Boolean]
  not: BooleanFilterInput
  is: Boolean
  null: Boolean
  notNull: Boolean
}

#input DateFilterInput {
#  and: [Date]
#  or: [Date]
#  not: DateFilterInput
#  eq: Date
#  ne: Date
#  gt: Date
#  gte: Date
#  lt: Date
#  lte: Date
#  null: Boolean
#  notNull: Boolean
#  in: [Date]
#  notIn: [Date]
#  between: DateFilterBetween
#}

#input DateFilterBetween{
#  start: Date!
#  end: Date!
#}
{{- range $objectName, $object := $root.List.Objects}}

input {{$object.Name}}Input{
  {{- range $entityKey, $entity := $object.Entities}}
  {{$entity.Name}}: {{$entity.GqlType "Input"}}{{$entity.RequiredChar}}
  {{- end}}
}

input {{$object.Name}}Patch{
  {{- range $entityKey, $entity := $object.Entities}}
  {{$entity.Name}}: {{$entity.GqlType "Patch"}}
  {{- end}}
} 


{{- if $object.HasSqlDirective}}

input Update{{$object.Name}}Input{
  filter: {{$object.Name}}FiltersInput
  set: {{$object.Name}}Patch
}

type Add{{$object.Name}}Payload{
  {{lcFirst $object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int): {{$object.Name}}QueryResult!
}

type Update{{$object.Name}}Payload{
  {{lcFirst  $object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int): {{$object.Name}}QueryResult!
  count: Int!
}

type Delete{{$object.Name}}Payload{
  {{lcFirst $object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int): {{$object.Name}}QueryResult!
  count: Int!
  msg: String
}

type {{$object.Name}}QueryResult{
  data: [{{$object.Name}}!]!
  count: Int!
  totalCount: Int!
}

enum {{$object.Name}}Orderable {
  {{- range $entityKey, $entity := $object.OrderAbleEntities}}
  {{$entity.Name}}
  {{- end}}
}

{{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
{{$refType := $root.List.PrimaryEntityOfObject $m2mEntity.GqlTypeName}}
input {{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput{
  filter: {{$object.Name}}FiltersInput!
  set: [{{$refType.GqlTypeName}}!]!
}
{{- end}}
input {{$object.Name}}Order{
  asc: {{$object.Name}}Orderable
  desc: {{$object.Name}}Orderable
}

input {{$object.Name}}FiltersInput{
  {{- range $entityKey, $entity := $object.Entities}}
  {{- if $entity.BuiltIn }}
  {{$entity.Name}}: {{$entity.GqlTypeName}}FilterInput
  {{- else}}
    {{- if $entity.GqlTypeObj.HasSqlDirective}}
  {{$entity.Name}}:{{$entity.GqlTypeName}}FiltersInput
    {{- end}}
  {{- end}}
  
  {{- end}}
  and: [{{$object.Name}}FiltersInput]
  or: [{{$object.Name}}FiltersInput]
  not: {{$object.Name}}FiltersInput
}


{{- if $object.SQLDirective.HasQueries}}
extend type Query {
    {{- if $object.SQLDirective.Query.Get}}
  get{{$object.Name}}({{range $entryKey, $entity := $object.PrimaryKeys}}{{$entity.Name}}: {{$entity.GqlType "Patch"}}{{end}}!): {{$object.Name}} {{ $object.SQLDirectiveValues "query" "Get" | join " "}}
    {{- end}}
    {{- if $object.SQLDirective.Query.Query}}
  query{{$object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int ): {{$object.Name}}QueryResult {{ $object.SQLDirectiveValues "query" "Query" | join " "}}
    {{- end}}
}
{{- end}}
{{- if $object.SQLDirective.HasMutation}}
extend type Mutation {
{{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
  add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s(input:{{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput!): Update{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Add" | join " "}}
{{- end}}
  {{- if $object.SQLDirective.Mutation.Add}}
  add{{$object.Name}}(input: [{{$object.Name}}Input!]!): Add{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Add" | join " "}}
  {{- end}}
  {{- if $object.SQLDirective.Mutation.Update}}
  update{{$object.Name}}(input: Update{{$object.Name}}Input!): Update{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Update" | join " "}}
  {{- end}}
  {{- if $object.SQLDirective.Mutation.Delete}}
  delete{{$object.Name}}(filter: {{$object.Name}}FiltersInput!): Delete{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Delete" | join " "}}
  {{- end}}
}
{{- end}}
{{- end}}

{{- end}}