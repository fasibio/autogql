// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package db

import (
	"context"

	"github.com/fasibio/autogql/testservice/graph/model"
	"gorm.io/gorm"
)

type QueryName string
type GetName string
type AddName string
type UpdateName string
type DeleteName string
type Many2ManyName string

const (
	GetCat           GetName       = "GetCat"
	QueryCat         QueryName     = "QueryCat"
	AddCat           AddName       = "AddCat"
	UpdateCat        UpdateName    = "UpdateCat"
	DeleteCat        DeleteName    = "DeleteCat"
	GetCompany       GetName       = "GetCompany"
	QueryCompany     QueryName     = "QueryCompany"
	AddCompany       AddName       = "AddCompany"
	UpdateCompany    UpdateName    = "UpdateCompany"
	DeleteCompany    DeleteName    = "DeleteCompany"
	GetSmartPhone    GetName       = "GetSmartPhone"
	QuerySmartPhone  QueryName     = "QuerySmartPhone"
	AddSmartPhone    AddName       = "AddSmartPhone"
	UpdateSmartPhone UpdateName    = "UpdateSmartPhone"
	DeleteSmartPhone DeleteName    = "DeleteSmartPhone"
	AddUser2Todos    Many2ManyName = "AddUser2Todos"
	GetTodo          GetName       = "GetTodo"
	QueryTodo        QueryName     = "QueryTodo"
	AddTodo          AddName       = "AddTodo"
	UpdateTodo       UpdateName    = "UpdateTodo"
	DeleteTodo       DeleteName    = "DeleteTodo"
	GetUser          GetName       = "GetUser"
	QueryUser        QueryName     = "QueryUser"
	AddUser          AddName       = "AddUser"
	UpdateUser       UpdateName    = "UpdateUser"
	DeleteUser       DeleteName    = "DeleteUser"
)

type AutoGqlHookM interface {
	model.Cat | model.Company | model.SmartPhone | model.Todo | model.User
}
type AutoGqlHookF interface {
	model.CatFiltersInput | model.CompanyFiltersInput | model.SmartPhoneFiltersInput | model.TodoFiltersInput | model.UserFiltersInput
}
type AutoGqlHookM2M interface {
	model.UserRef2TodosInput
}

type AutoGqlHookQueryO interface {
	model.CatOrder | model.CompanyOrder | model.SmartPhoneOrder | model.TodoOrder | model.UserOrder
}

type AutoGqlHookI interface {
	model.CatInput | model.CompanyInput | model.SmartPhoneInput | model.TodoInput | model.UserInput
}

type AutoGqlHookU interface {
	model.UpdateCatInput | model.UpdateCompanyInput | model.UpdateSmartPhoneInput | model.UpdateTodoInput | model.UpdateUserInput
}

type AutoGqlHookUP interface {
	model.UpdateCatPayload | model.UpdateCompanyPayload | model.UpdateSmartPhonePayload | model.UpdateTodoPayload | model.UpdateUserPayload
}

type AutoGqlHookDP interface {
	model.DeleteCatPayload | model.DeleteCompanyPayload | model.DeleteSmartPhonePayload | model.DeleteTodoPayload | model.DeleteUserPayload
}

type AutoGqlHookAP interface {
	model.AddCatPayload | model.AddCompanyPayload | model.AddSmartPhonePayload | model.AddTodoPayload | model.AddUserPayload
}

func AddGetHook[T AutoGqlHookM, I any](db *AutoGqlDB, name GetName, implementation AutoGqlHookGet[T, I]) {
	db.Hooks[string(name)] = implementation
}

func AddQueryHook[M AutoGqlHookM, F AutoGqlHookF, O AutoGqlHookQueryO](db *AutoGqlDB, name QueryName, implementation AutoGqlHookQuery[M, F, O]) {
	db.Hooks[string(name)] = implementation
}

func AddAddHook[M AutoGqlHookM, I AutoGqlHookI, AP AutoGqlHookAP](db *AutoGqlDB, name AddName, implementation AutoGqlHookAdd[M, I, AP]) {
	db.Hooks[string(name)] = implementation
}

func AddUpdateHook[M AutoGqlHookM, U AutoGqlHookU, UP AutoGqlHookUP](db *AutoGqlDB, name UpdateName, implementation AutoGqlHookUpdate[U, UP]) {
	db.Hooks[string(name)] = implementation
}

func AddMany2ManyHook[U AutoGqlHookM2M, UP AutoGqlHookUP](db *AutoGqlDB, name Many2ManyName, implementation AutoGqlHookMany2Many[U, UP]) {
	db.Hooks[string(name)] = implementation
}

func AddDeleteHook[F AutoGqlHookF, DP AutoGqlHookDP](db *AutoGqlDB, name DeleteName, implementation AutoGqlHookDelete[F, DP]) {
	db.Hooks[string(name)] = implementation
}

type AutoGqlHookGet[obj AutoGqlHookM, identifier any] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, id ...identifier) (*gorm.DB, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data *obj) (*obj, error)
	BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error)
}

type DefaultGetHook[obj AutoGqlHookM, identifier any] struct{}

func (d DefaultGetHook[obj, identifier]) Received(ctx context.Context, dbHelper *AutoGqlDB, id ...identifier) (*gorm.DB, error) {
	return dbHelper.Db, nil
}
func (d DefaultGetHook[obj, identifier]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}
func (d DefaultGetHook[obj, identifier]) AfterCallDb(ctx context.Context, data *obj) (*obj, error) {
	return data, nil
}
func (d DefaultGetHook[obj, identifier]) BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error) {
	return data, nil
}

type AutoGqlHookQuery[obj AutoGqlHookM, filter AutoGqlHookF, order AutoGqlHookQueryO] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, filter *filter, order *order, first, offset *int) (*gorm.DB, *filter, *order, *int, *int, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error)
	BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error)
}

type DefaultQueryHook[obj AutoGqlHookM, filter AutoGqlHookF, order AutoGqlHookQueryO] struct{}

func (d DefaultQueryHook[obj, filterType, orderType]) Received(ctx context.Context, dbHelper *AutoGqlDB, filter *filterType, order *orderType, first, offset *int) (*gorm.DB, *filterType, *orderType, *int, *int, error) {
	return dbHelper.Db, filter, order, first, offset, nil
}
func (d DefaultQueryHook[obj, filter, order]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}
func (d DefaultQueryHook[obj, filter, order]) AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error) {
	return data, nil
}
func (d DefaultQueryHook[obj, filter, order]) BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error) {
	return data, nil
}

type AutoGqlHookAdd[obj AutoGqlHookM, input AutoGqlHookI, res AutoGqlHookAP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input []*input) (*gorm.DB, []*input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB, []obj, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, data []obj, res *res) (*res, error)
}

type DefaultAddHook[obj AutoGqlHookM, input AutoGqlHookI, res AutoGqlHookAP] struct{}

func (d DefaultAddHook[obj, inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input []*inputType) (*gorm.DB, []*inputType, error) {
	return dbHelper.Db, input, nil
}
func (d DefaultAddHook[obj, inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB, []obj, error) {
	return db, data, nil
}
func (d DefaultAddHook[obj, inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, data []obj, res *resType) (*resType, error) {
	return res, nil
}

type AutoGqlHookUpdate[input AutoGqlHookU, res AutoGqlHookUP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type AutoGqlHookMany2Many[input AutoGqlHookM2M, res AutoGqlHookUP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type DefaultMany2ManyHook[input AutoGqlHookM2M, res AutoGqlHookUP] struct{}

func (d DefaultMany2ManyHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}
func (d DefaultMany2ManyHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}
func (d DefaultMany2ManyHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}

type DefaultUpdateHook[input AutoGqlHookU, res AutoGqlHookUP] struct{}

func (d DefaultUpdateHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}
func (d DefaultUpdateHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error) {
	return db, data, nil
}
func (d DefaultUpdateHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}

type AutoGqlHookDelete[input AutoGqlHookF, res AutoGqlHookDP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type DefaultDeleteHook[input AutoGqlHookF, res AutoGqlHookDP] struct{}

func (d DefaultDeleteHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}
func (d DefaultDeleteHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}
func (d DefaultDeleteHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}
