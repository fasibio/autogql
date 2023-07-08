package testservice

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fasibio/autogql/testservice/graph/model"
)

func ValidateDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	field := graphql.GetPathContext(ctx)

	if data, ok := obj.(map[string]interface{}); ok {
		for _, v := range field.ParentField.Field.Arguments {
			model, err := model.GetInputStruct(v.Value.ExpectedType.Name(), data)
			if err != nil {
				continue
			}
			if err := validate.Struct(model); err != nil {
				return nil, fmt.Errorf("validationerror:  %v", err)
			}
		}
	}
	return next(ctx)
}
