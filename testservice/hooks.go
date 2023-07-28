package testservice

import (
	"context"

	"github.com/fasibio/autogql/testservice/graph/db"
	"github.com/fasibio/autogql/testservice/graph/model"
	"gorm.io/gorm"
)

type AddTodoHook struct {
	db.DefaultAddHook[model.Todo, model.TodoInput, model.AddTodoPayload]
}

func (a AddTodoHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Todo) (*gorm.DB, []model.Todo, error) {
	newData := make([]model.Todo, len(data))
	for i, d := range data {
		d.OwnerID = 6
		newData[i] = d

	}
	return db, newData, nil
}
