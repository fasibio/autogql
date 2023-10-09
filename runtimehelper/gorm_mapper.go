package runtimehelper

import (
	"fmt"
)

type Relation string

const (
	RelationAnd Relation = "AND"
	RelationOr  Relation = "OR"
	RelationNot Relation = "NOT"
)

type ConditionElement struct {
	Operator         string             `json:"operator,omitempty"`
	Field            string             `json:"field,omitempty"`
	Value            []interface{}      `json:"value,omitempty"`
	Children         []ConditionElement `json:"children,omitempty"`
	ChildrenRelation Relation           `json:"children_relation,omitempty"`
}

func CombineSimpleQuery(elements []ConditionElement, relation Relation) (string, []interface{}) {
	if len(elements) == 0 {
		return "", nil
	}
	sql := ""
	values := make([]interface{}, 0, len(elements))
	for _, query := range elements {
		if sql == "" {
			if len(query.Children) == 0 {
				if len(query.Value) == 0 {
					sql += fmt.Sprintf("%s %s", query.Field, query.Operator)
				} else {
					sql += fmt.Sprintf("%s %s ?", query.Field, query.Operator)
					values = append(values, query.Value...)
				}
			} else {
				querySql, queryValues := CombineSimpleQuery(query.Children, query.ChildrenRelation)
				sql += fmt.Sprintf("(%s)", querySql)
				values = append(values, queryValues...)
			}
		} else {
			if len(query.Children) == 0 {
				if len(query.Value) == 0 {
					sql += fmt.Sprintf(" %s %s %s", relation, query.Field, query.Operator)
				} else {
					sql += fmt.Sprintf(" %s %s %s ?", relation, query.Field, query.Operator)
					values = append(values, query.Value...)
				}
			} else {
				querySql, queryValues := CombineSimpleQuery(query.Children, query.ChildrenRelation)
				sql += fmt.Sprintf(" %s (%s)", relation, querySql)
				values = append(values, queryValues...)
			}
		}
	}
	return sql, values
}

func Complex(relation Relation, elems ...ConditionElement) ConditionElement {
	return ConditionElement{
		Children:         elems,
		ChildrenRelation: relation,
	}
}

func Equal(field string, value interface{}) ConditionElement {
	return NewConditionElement("=", field, value)
}

func NotEqual(field string, value interface{}) ConditionElement {
	return NewConditionElement("<>", field, value)
}

func Like(field string, value interface{}) ConditionElement {
	return NewConditionElement("LIKE", field, value)
}

func NotLike(field string, value interface{}) ConditionElement {
	return NewConditionElement("NOT LIKE", field, value)
}

func More(field string, value interface{}) ConditionElement {
	return NewConditionElement(">", field, value)
}

func MoreOrEqual(field string, value interface{}) ConditionElement {
	return NewConditionElement(">=", field, value)
}

func Less(field string, value interface{}) ConditionElement {
	return NewConditionElement("<", field, value)
}

func LessOrEqual(field string, value interface{}) ConditionElement {
	return NewConditionElement("<=", field, value)
}

func In(field string, value interface{}) ConditionElement {
	return NewConditionElement("IN", field, value)
}

func NotIn(field string, value interface{}) ConditionElement {
	return NewConditionElement("NOT IN", field, value)
}

func NotNull(field string) ConditionElement {
	return NewConditionElement("IS NOT NULL", field)
}

func Null(field string) ConditionElement {
	return NewConditionElement("IS NULL", field)
}

func Between(field string, start, end interface{}) ConditionElement {
	return NewConditionElement("BETWEEN ? AND", field, start, end)
}

func NewConditionElement(operator, field string, value ...interface{}) ConditionElement {
	return ConditionElement{
		Operator: operator,
		Field:    field,
		Value:    value,
	}
}
