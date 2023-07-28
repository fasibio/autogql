package testservice

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fasibio/autogql/testservice/graph/model"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
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
				switch err := err.(type) {
				case validator.ValidationErrors:
					{
						errList := gqlerror.List{}
						for _, v := range err {

							errList = append(errList, &gqlerror.Error{
								Path:    graphql.GetPath(ctx),
								Message: v.Error(),
								Extensions: map[string]interface{}{
									"code":      "INPUT_VALIDATION_ERROR",
									"tag":       v.Tag(),
									"field":     v.Field(),
									"namespace": v.Namespace(),
								},
							})
						}
						return nil, errList
					}
				default:
					{
						return nil, fmt.Errorf("validationerror:  %v", err)
					}
				}

			}
		}
	}
	return next(ctx)
}
