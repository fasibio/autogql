{{- $root := .}}
{{- $commentFilterText := "Filter simple datatypes"}}
"""
ID {{ $commentFilterText }}
"""
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

"""
String {{ $commentFilterText }}
"""
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

"""
Int {{ $commentFilterText }}
"""
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

"""
Filter between start and end (start > value < end)
"""
input IntFilterBetween{
  start: Int!
  end: Int!
}

"""
Boolean {{ $commentFilterText }}
"""
input BooleanFilterInput{
  and: [Boolean]
  or: [Boolean]
  not: BooleanFilterInput
  is: Boolean
  null: Boolean
  notNull: Boolean
}

"""
Time {{ $commentFilterText }}
"""
input TimeFilterInput {
  and: [Time]
  or: [Time]
  not: TimeFilterInput
  eq: Time
  ne: Time
  gt: Time
  gte: Time
  lt: Time
  lte: Time
  null: Boolean
  notNull: Boolean
  in: [Time]
  notIn: [Time]
  between: TimeFilterBetween
}

"""
Filter between start and end (start > value < end)
"""
input TimeFilterBetween{
  start: Time!
  end: Time!
}
{{- range $objectName, $object := $root.List.Objects}}
  {{- if $object.HasSqlDirective}}

  """
  {{$object.Name}} Input value to add new {{$object.Name}}
  """
  input {{$object.Name}}Input {{$object.InputTypeDirectiveGql}}{
    {{- range $entityKey, $entity := $object.InputEntities}}
      {{$entity.Name}}: {{$entity.GqlType "Input"}}{{$entity.RequiredChar}} {{$entity.InputTypeDirectiveGql}}
    {{- end}}
  }

  """
  {{$object.Name}} Patch value all values are optional to update {{$object.Name}} entities
  """
  input {{$object.Name}}Patch {{$object.InputTypeDirectiveGql}}{
    {{- range $entityKey, $entity := $object.PatchEntities}}
      {{$entity.Name}}: {{$entity.GqlType "Patch"}} {{$entity.InputTypeDirectiveGql}}
    {{- end}}
  } 


    """
    Update rules for {{$object.Name}} multiupdates simple possible by global filtervalue
    """
    input Update{{$object.Name}}Input{
      filter: {{$object.Name}}FiltersInput!
      set: {{$object.Name}}Patch!
    }

    """
    Add{{$object.Name}} result with filterable data and affected rows
    """
    type Add{{$object.Name}}Payload{
      {{lcFirst $object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int, group: [{{$object.Name}}Group!]): {{$object.Name}}QueryResult!
      affected: [{{$object.Name}}!]!
    }

    """
    Update{{$object.Name}} result with filterable data and affected rows
    """
    type Update{{$object.Name}}Payload{
      {{lcFirst  $object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int, group: [{{$object.Name}}Group!]): {{$object.Name}}QueryResult!
      """
      Count of affected updates
      """
      count: Int!
      affected: [{{$object.Name}}!]!
    }

    """
    Delete{{$object.Name}} result with filterable data and count of affected entries
    """
    type Delete{{$object.Name}}Payload{
      {{lcFirst $object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int, group: [{{$object.Name}}Group!]): {{$object.Name}}QueryResult!
      """
      Count of deleted {{$object.Name}} entities
      """
      count: Int!
      msg: String
    }

    """
    {{$object.Name}} result
    """
    type {{$object.Name}}QueryResult{
      data: [{{$object.Name}}!]!
      count: Int!
      totalCount: Int!
    }

    """
    for {{$object.Name}} a enum of all orderable entities
    can be used f.e.: query{{$object.Name}}
    """
    enum {{$object.Name}}Orderable {
      {{- range $entityKey, $entity := $object.OrderAbleEntities}}
        {{$entity.Name}}
      {{- end}}
    }

    {{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
      {{$refType := $root.List.PrimaryEntityOfObject $m2mEntity.GqlTypeName}}
      """
      Many 2 many input between {{$object.Name}} and {{$m2mEntity.GqlTypeName}}
      Filter to Select {{$object.Name}} and set to set list of {{$m2mEntity.GqlTypeName}} PrimaryKeys
      """
      input {{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput{
        filter: {{$object.Name}}FiltersInput!
        set: [{{$refType.GqlTypeName}}!]!
      }
    {{- end}}
    """
    Order {{$object.Name}} by asc or desc 
    """
    input {{$object.Name}}Order{
      asc: {{$object.Name}}Orderable
      desc: {{$object.Name}}Orderable
    }

    """
    Groupable data for  {{$object.Name}}
    Can be used f.e.: by query{{$object.Name}}
    """
    enum {{$object.Name}}Group {
      {{- range $entityKey, $entity := $object.InputFilterEntities}}
        {{- if $entity.IsPrimitive }}
          {{$entity.Name}}
        {{- end }}
      {{- end}}
    }

    """
    Filter input selection for {{$object.Name}}
    Can be used f.e.: by query{{$object.Name}}
    """
    input {{$object.Name}}FiltersInput{
      {{- range $entityKey, $entity := $object.InputFilterEntities}}
        {{- if $entity.IsPrimitive }}
          {{$entity.Name}}: {{$entity.GqlTypeName}}FilterInput
        {{- else}}
          {{- if $entity.IsEnum}}
            {{$entity.Name}}: StringFilterInput
          {{- else }}
            {{- if $entity.GqlTypeObj.HasSqlDirective}}
              {{$entity.Name}}:{{$entity.GqlTypeName}}FiltersInput
            {{- end}}
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
        """
        return one {{$object.Name}} selected by PrimaryKey(s)
        """
        get{{$object.Name}}({{range $entryKey, $entity := $object.PrimaryKeys}}{{$entity.Name}}: {{$entity.GqlType "Patch"}}!, {{end}}): {{$object.Name}} {{ $object.SQLDirectiveValues "query" "Get" | join " "}}
      {{- end}}
      {{- if $object.SQLDirective.Query.Query}}
        """
        return a list of  {{$object.Name}} filterable, pageination, orderbale, groupable ...
        """
        query{{$object.Name}}(filter: {{$object.Name}}FiltersInput, order: {{$object.Name}}Order, first: Int, offset: Int, group: [{{$object.Name}}Group!] ): {{$object.Name}}QueryResult {{ $object.SQLDirectiveValues "query" "Query" | join " "}}
      {{- end}}
      }
    {{- end}}
    {{- if $object.SQLDirective.HasMutation}}
      extend type Mutation {
      {{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
        """
        Add new Many2Many relation(s)
        """
        add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s(input:{{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput!): Update{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Add" | join " "}}

        """
        Delete Many2Many relation(s)
        """
        delete{{$m2mEntity.GqlTypeName}}From{{$object.Name}}s(input:{{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput!): Delete{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Delete" | join " "}}
      {{- end}}
      {{- if $object.SQLDirective.Mutation.Add}}
        """
        Add new {{$object.Name}}
        """
        add{{$object.Name}}(input: [{{$object.Name}}Input!]!): Add{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Add" | join " "}}
      {{- end}}
      {{- if $object.SQLDirective.Mutation.Update}}
        """
        update {{$object.Name}} filtered by selection and update all matched values
        """
        update{{$object.Name}}(input: Update{{$object.Name}}Input!): Update{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Update" | join " "}}
      {{- end}}
      {{- if $object.SQLDirective.Mutation.Delete}}
        """
        delete {{$object.Name}} filtered by selection and delete all matched values
        """
        delete{{$object.Name}}(filter: {{$object.Name}}FiltersInput!): Delete{{$object.Name}}Payload {{ $object.SQLDirectiveValues "mutation" "Delete" | join " "}}
      {{- end}}
      }
    {{- end}}
  {{- end}}
{{- end}}