// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	"github.com/fasibio/autogql/runtimehelper"
	"gorm.io/gorm"
)

type ParentObject interface {
	TableName() string
	PrimaryKeyName() string
}

func (d *CatFiltersInput) TableName() string {
	return "cat"
}

func (d *CatFiltersInput) PrimaryKeyName() string {
	return "id"
}

func (d *CatFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.BirthDay != nil {
		res = append(res, d.BirthDay.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "birth_day"), true, blackList)...)
	}
	if d.UserID != nil {
		res = append(res, d.UserID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "user_id"), true, blackList)...)
	}
	if d.Alive != nil {
		res = append(res, d.Alive.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "alive"), true, blackList)...)
	}

	return res
}

func (d *CompanyFiltersInput) TableName() string {
	return "company"
}

func (d *CompanyFiltersInput) PrimaryKeyName() string {
	return "id"
}

func (d *CompanyFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.Description != nil {
		res = append(res, d.Description.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "description"), true, blackList)...)
	}
	if d.MotherCompanyID != nil {
		res = append(res, d.MotherCompanyID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "mother_company_id"), true, blackList)...)
	}
	if d.MotherCompany != nil {
		res = append(res, d.MotherCompany.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]sMotherCompany%[1]s", runtimehelper.GetQuoteChar(db)), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}

	return res
}

func (d *TodoFiltersInput) TableName() string {
	return "todo"
}

func (d *TodoFiltersInput) PrimaryKeyName() string {
	return "id"
}

func (d *TodoFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.Users != nil {
		tableName := db.Config.NamingStrategy.TableName("User")
		if _, ok := blackList["todo_users"]; !ok {
			blackList["todo_users"] = struct{}{}
			db = db.Joins(fmt.Sprintf("LEFT JOIN %[1]stodo_users%[1]s ON %[1]stodo_users%[1]s.%[1]stodo_id%[1]s = %[2]s.%[1]sid%[1]s JOIN %[1]s%[3]s%[1]s ON %[1]stodo_users%[1]s.%[1]suser_id%[1]s = %[1]s%[3]s%[1]s.%[1]sid%[1]s", runtimehelper.GetQuoteChar(db), alias, tableName))
		}
		res = append(res, d.Users.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), true, blackList)...)
	}
	if d.Owner != nil {
		if _, ok := blackList["Owner"]; !ok {
			blackList["Owner"] = struct{}{}
			if deep {
				tableName := db.Config.NamingStrategy.TableName("User")
				foreignKeyName := "owner_id"
				db = db.Joins(fmt.Sprintf("LEFT JOIN %[1]s%[2]s%[1]s %[1]sOwner%[1]s ON %[1]sOwner%[1]s.%[1]s%[3]s%[1]s = %[4]s.%[1]s%[5]s%[1]s", runtimehelper.GetQuoteChar(db), tableName, d.PrimaryKeyName(), alias, foreignKeyName))

			} else {
				db = db.Joins("Owner")
			}
		}
		res = append(res, d.Owner.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]sOwner%[1]s", runtimehelper.GetQuoteChar(db)), true, blackList)...)
	}
	if d.OwnerID != nil {
		res = append(res, d.OwnerID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "owner_id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.DeletedAt != nil {
		res = append(res, d.DeletedAt.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "deleted_at"), true, blackList)...)
	}

	return res
}

func (d *UserFiltersInput) TableName() string {
	return "user"
}

func (d *UserFiltersInput) PrimaryKeyName() string {
	return "id"
}

func (d *UserFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.DeletedAt != nil {
		res = append(res, d.DeletedAt.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "deleted_at"), true, blackList)...)
	}
	if d.Cat != nil {
		if _, ok := blackList["Cat"]; !ok {
			blackList["Cat"] = struct{}{}
			if deep {
				tableName := db.Config.NamingStrategy.TableName("Cat")
				foreignKeyName := "user_id"
				db = db.Joins(fmt.Sprintf("LEFT JOIN %[1]s%[2]s%[1]s %[1]sCat%[1]s ON %[1]sCat%[1]s.%[1]s%[5]s%[1]s = %[4]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), tableName, d.PrimaryKeyName(), alias, foreignKeyName))

			} else {
				db = db.Joins("Cat")
			}
		}
		res = append(res, d.Cat.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]sCat%[1]s", runtimehelper.GetQuoteChar(db)), true, blackList)...)
	}
	if d.CompanyID != nil {
		res = append(res, d.CompanyID.ExtendsDatabaseQuery(db, fmt.Sprintf("%[2]s.%[1]s%[3]s%[1]s", runtimehelper.GetQuoteChar(db), alias, "company_id"), true, blackList)...)
	}
	if d.Company != nil {
		if _, ok := blackList["Company"]; !ok {
			blackList["Company"] = struct{}{}
			if deep {
				tableName := db.Config.NamingStrategy.TableName("Company")
				foreignKeyName := "company_id"
				db = db.Joins(fmt.Sprintf("LEFT JOIN %[1]s%[2]s%[1]s %[1]sCompany%[1]s ON %[1]sCompany%[1]s.%[1]s%[3]s%[1]s = %[4]s.%[1]s%[5]s%[1]s", runtimehelper.GetQuoteChar(db), tableName, d.PrimaryKeyName(), alias, foreignKeyName))

			} else {
				db = db.Joins("Company")
			}
		}
		res = append(res, d.Company.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]sCompany%[1]s", runtimehelper.GetQuoteChar(db)), true, blackList)...)
	}

	return res
}

func (d *StringFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}
	if d.Contains != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%%%s%%", *d.Contains)))
	}

	if d.Containsi != nil {
		res = append(res, runtimehelper.Like(fmt.Sprintf("lower(%s)", fieldName), fmt.Sprintf("%%%s%%", strings.ToLower(*d.Containsi))))
	}

	if d.EndsWith != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%%%s", *d.EndsWith)))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}

	if d.Eqi != nil {
		res = append(res, runtimehelper.Equal(fmt.Sprintf("lower(%s)", fieldName), strings.ToLower(*d.Eqi)))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotContains != nil {
		res = append(res, runtimehelper.NotLike(fieldName, fmt.Sprintf("%%%s%%", *d.NotContains)))
	}

	if d.NotContainsi != nil {
		res = append(res, runtimehelper.NotLike(fmt.Sprintf("lower(%s)", fieldName), fmt.Sprintf("%%%s%%", strings.ToLower(*d.NotContainsi))))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))
	}

	if d.NotNull != nil {
		res = append(res, runtimehelper.NotNull(fieldName, d.NotNull))
	}

	if d.Null != nil {
		res = append(res, runtimehelper.Null(fieldName, d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.StartsWith != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%s%%", *d.StartsWith)))
	}

	return res
}

func (d *IntFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName, d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName, *d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName, *d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName, *d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName, *d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))

	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName, *d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName, *d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

func (d *BooleanFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Is != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Is))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName, *d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName, *d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

func (d *TimeFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName, d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName, *d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName, *d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName, *d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName, *d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))

	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName, *d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName, *d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

func (d *IDFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName, *d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName, *d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}