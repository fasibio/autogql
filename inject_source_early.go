package autogql

import (
	"fmt"
	"log"

	"github.com/fasibio/autogql/structure"
	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs *AutoGqlPlugin) InjectSourceEarly() *ast.Source {
	log.Println("InjectSourceEarly")

	input := fmt.Sprintf(`

	input SqlCreateExtension {
		value: Boolean!
		directiveExt: [String!]
	}

	input SqlMutationParams {
		add: SqlCreateExtension
		update: SqlCreateExtension
		delete: SqlCreateExtension
		directiveExt: [String!]
	}

	input SqlQueryParams {
		get: SqlCreateExtension
		query: SqlCreateExtension
		directiveExt: [String!]
	}
	directive @%s(order: Int, %s:SqlQueryParams, %s: SqlMutationParams ) on OBJECT
	directive @%s on FIELD_DEFINITION
	directive @%s on FIELD_DEFINITION

	directive @%s (value: String)on FIELD_DEFINITION
  
	directive @%s on FIELD_DEFINITION

	directive @%s (value: [String!]) on FIELD_DEFINITION

	directive @%s (value: [String!]) on FIELD_DEFINITION | OBJECT

	scalar Time
	scalar SoftDelete

# internal directive
directive @%s (value: [String!]) on INPUT_FIELD_DEFINITION

## internal directive end
	`, structure.DirectiveSQL,
		structure.DirectiveSQLArgumentQuery,
		structure.DirectiveSQLArgumentMutation,
		structure.DirectiveSQLPrimary,
		structure.DirectiveSQLIndex,
		structure.DirectiveSQLGorm,
		structure.DirectiveNoMutation,
		structure.DirectiveSQLInputTypeTags,
		structure.DirectiveSQLInputTypeDirective,
		structure.DirectiveSQLInputTypeTags.InternalName())

	return &ast.Source{
		Name:    fmt.Sprintf("%s/directive.graphql", ggs.Name()),
		Input:   input,
		BuiltIn: true,
	}
}
