package model

import (
	"context"
	"time"
)

type Todo struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name      string     `json:"name"`
	Users     []*User    `json:"users" gorm:"many2many:todo_users;constraint:OnDelete:CASCADE;"`
	Owner     *User      `json:"owner"`
	OwnerID   int        `json:"ownerID"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Etype1    *TodoType  `json:"etype1,omitempty"`
	Etype5    TodoType   `json:"etype5"`
	Test123   *string    `json:"test123,omitempty"`
}

func (t *Todo) NoControl(ctx context.Context) *NoSQLControl {
	ares := "AAAA"
	return &NoSQLControl{
		ID: 1234,
		A:  &ares,
		B:  2222,
	}
}
