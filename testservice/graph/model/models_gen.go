// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AddCatPayload struct {
	Cat *CatQueryResult `json:"cat"`
}

type AddCompanyPayload struct {
	Company *CompanyQueryResult `json:"company"`
}

type AddSmartPhonePayload struct {
	SmartPhone *SmartPhoneQueryResult `json:"smartPhone"`
}

type AddTodoPayload struct {
	Todo *TodoQueryResult `json:"todo"`
}

type AddUserPayload struct {
	User *UserQueryResult `json:"user"`
}

type BooleanFilterInput struct {
	And     []*bool             `json:"and,omitempty"`
	Or      []*bool             `json:"or,omitempty"`
	Not     *BooleanFilterInput `json:"not,omitempty"`
	Is      *bool               `json:"is,omitempty"`
	Null    *bool               `json:"null,omitempty"`
	NotNull *bool               `json:"notNull,omitempty"`
}

type Cat struct {
	ID       int       `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birthDay"`
	Age      *int      `json:"age,omitempty" gorm:"-;"`
	UserID   int       `json:"userID"`
	Alive    *bool     `json:"alive,omitempty" gorm:"default:true;"`
}

type CatFiltersInput struct {
	ID       *IDFilterInput      `json:"id,omitempty"`
	Name     *StringFilterInput  `json:"name,omitempty"`
	BirthDay *TimeFilterInput    `json:"birthDay,omitempty"`
	UserID   *IntFilterInput     `json:"userID,omitempty"`
	Alive    *BooleanFilterInput `json:"alive,omitempty"`
	And      []*CatFiltersInput  `json:"and,omitempty"`
	Or       []*CatFiltersInput  `json:"or,omitempty"`
	Not      *CatFiltersInput    `json:"not,omitempty"`
}

type CatInput struct {
	Name     string    `json:"name"`
	BirthDay time.Time `json:"birthDay"`
	UserID   int       `json:"userID"`
	Alive    *bool     `json:"alive,omitempty"`
}

type CatOrder struct {
	Asc  *CatOrderable `json:"asc,omitempty"`
	Desc *CatOrderable `json:"desc,omitempty"`
}

type CatPatch struct {
	Name     *string    `json:"name,omitempty"`
	BirthDay *time.Time `json:"birthDay,omitempty"`
	UserID   *int       `json:"userID,omitempty"`
	Alive    *bool      `json:"alive,omitempty"`
}

type CatQueryResult struct {
	Data       []*Cat `json:"data"`
	Count      int    `json:"count"`
	TotalCount int    `json:"totalCount"`
}

type Company struct {
	ID              int        `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name            string     `json:"name"`
	Description     *string    `json:"description,omitempty"`
	MotherCompanyID *int       `json:"motherCompanyID,omitempty"`
	MotherCompany   *Company   `json:"motherCompany,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
}

type CompanyFiltersInput struct {
	ID              *IDFilterInput         `json:"id,omitempty"`
	Name            *StringFilterInput     `json:"name,omitempty"`
	Description     *StringFilterInput     `json:"description,omitempty"`
	MotherCompanyID *IntFilterInput        `json:"motherCompanyID,omitempty"`
	MotherCompany   *CompanyFiltersInput   `json:"motherCompany,omitempty"`
	CreatedAt       *TimeFilterInput       `json:"createdAt,omitempty"`
	And             []*CompanyFiltersInput `json:"and,omitempty"`
	Or              []*CompanyFiltersInput `json:"or,omitempty"`
	Not             *CompanyFiltersInput   `json:"not,omitempty"`
}

type CompanyInput struct {
	Name            string        `json:"name"`
	Description     *string       `json:"description,omitempty"`
	MotherCompanyID *int          `json:"motherCompanyID,omitempty"`
	MotherCompany   *CompanyInput `json:"motherCompany,omitempty"`
}

type CompanyOrder struct {
	Asc  *CompanyOrderable `json:"asc,omitempty"`
	Desc *CompanyOrderable `json:"desc,omitempty"`
}

type CompanyPatch struct {
	Name            *string       `json:"name,omitempty"`
	Description     *string       `json:"description,omitempty"`
	MotherCompanyID *int          `json:"motherCompanyID,omitempty"`
	MotherCompany   *CompanyPatch `json:"motherCompany,omitempty"`
}

type CompanyQueryResult struct {
	Data       []*Company `json:"data"`
	Count      int        `json:"count"`
	TotalCount int        `json:"totalCount"`
}

type DeleteCatPayload struct {
	Cat   *CatQueryResult `json:"cat"`
	Count int             `json:"count"`
	Msg   *string         `json:"msg,omitempty"`
}

type DeleteCompanyPayload struct {
	Company *CompanyQueryResult `json:"company"`
	Count   int                 `json:"count"`
	Msg     *string             `json:"msg,omitempty"`
}

type DeleteSmartPhonePayload struct {
	SmartPhone *SmartPhoneQueryResult `json:"smartPhone"`
	Count      int                    `json:"count"`
	Msg        *string                `json:"msg,omitempty"`
}

type DeleteTodoPayload struct {
	Todo  *TodoQueryResult `json:"todo"`
	Count int              `json:"count"`
	Msg   *string          `json:"msg,omitempty"`
}

type DeleteUserPayload struct {
	User  *UserQueryResult `json:"user"`
	Count int              `json:"count"`
	Msg   *string          `json:"msg,omitempty"`
}

type IDFilterInput struct {
	And     []*int         `json:"and,omitempty"`
	Or      []*int         `json:"or,omitempty"`
	Not     *IDFilterInput `json:"not,omitempty"`
	Eq      *int           `json:"eq,omitempty"`
	Ne      *int           `json:"ne,omitempty"`
	Null    *bool          `json:"null,omitempty"`
	NotNull *bool          `json:"notNull,omitempty"`
	In      []*int         `json:"in,omitempty"`
	Notin   []*int         `json:"notin,omitempty"`
}

type IntFilterBetween struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type IntFilterInput struct {
	And     []*int            `json:"and,omitempty"`
	Or      []*int            `json:"or,omitempty"`
	Not     *IntFilterInput   `json:"not,omitempty"`
	Eq      *int              `json:"eq,omitempty"`
	Ne      *int              `json:"ne,omitempty"`
	Gt      *int              `json:"gt,omitempty"`
	Gte     *int              `json:"gte,omitempty"`
	Lt      *int              `json:"lt,omitempty"`
	Lte     *int              `json:"lte,omitempty"`
	Null    *bool             `json:"null,omitempty"`
	NotNull *bool             `json:"notNull,omitempty"`
	In      []*int            `json:"in,omitempty"`
	NotIn   []*int            `json:"notIn,omitempty"`
	Between *IntFilterBetween `json:"between,omitempty"`
}

type SmartPhone struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement;"`
	Brand       string `json:"brand"`
	Phonenumber string `json:"phonenumber"`
	UserID      int    `json:"userID"`
}

type SmartPhoneFiltersInput struct {
	ID          *IDFilterInput            `json:"id,omitempty"`
	Brand       *StringFilterInput        `json:"brand,omitempty"`
	Phonenumber *StringFilterInput        `json:"phonenumber,omitempty"`
	UserID      *IDFilterInput            `json:"userID,omitempty"`
	And         []*SmartPhoneFiltersInput `json:"and,omitempty"`
	Or          []*SmartPhoneFiltersInput `json:"or,omitempty"`
	Not         *SmartPhoneFiltersInput   `json:"not,omitempty"`
}

type SmartPhoneInput struct {
	Brand       string `json:"brand"`
	Phonenumber string `json:"phonenumber"`
	UserID      int    `json:"userID"`
}

type SmartPhoneOrder struct {
	Asc  *SmartPhoneOrderable `json:"asc,omitempty"`
	Desc *SmartPhoneOrderable `json:"desc,omitempty"`
}

type SmartPhonePatch struct {
	Brand       *string `json:"brand,omitempty"`
	Phonenumber *string `json:"phonenumber,omitempty"`
	UserID      *int    `json:"userID,omitempty"`
}

type SmartPhoneQueryResult struct {
	Data       []*SmartPhone `json:"data"`
	Count      int           `json:"count"`
	TotalCount int           `json:"totalCount"`
}

type SQLCreateExtension struct {
	Value        bool     `json:"value"`
	DirectiveExt []string `json:"directiveExt,omitempty"`
}

type SQLMutationParams struct {
	Add          *SQLCreateExtension `json:"add,omitempty"`
	Update       *SQLCreateExtension `json:"update,omitempty"`
	Delete       *SQLCreateExtension `json:"delete,omitempty"`
	DirectiveExt []string            `json:"directiveExt,omitempty"`
}

type SQLQueryParams struct {
	Get          *SQLCreateExtension `json:"get,omitempty"`
	Query        *SQLCreateExtension `json:"query,omitempty"`
	DirectiveExt []string            `json:"directiveExt,omitempty"`
}

type StringFilterInput struct {
	And          []*string          `json:"and,omitempty"`
	Or           []*string          `json:"or,omitempty"`
	Not          *StringFilterInput `json:"not,omitempty"`
	Eq           *string            `json:"eq,omitempty"`
	Eqi          *string            `json:"eqi,omitempty"`
	Ne           *string            `json:"ne,omitempty"`
	StartsWith   *string            `json:"startsWith,omitempty"`
	EndsWith     *string            `json:"endsWith,omitempty"`
	Contains     *string            `json:"contains,omitempty"`
	NotContains  *string            `json:"notContains,omitempty"`
	Containsi    *string            `json:"containsi,omitempty"`
	NotContainsi *string            `json:"notContainsi,omitempty"`
	Null         *bool              `json:"null,omitempty"`
	NotNull      *bool              `json:"notNull,omitempty"`
	In           []*string          `json:"in,omitempty"`
	NotIn        []*string          `json:"notIn,omitempty"`
}

type TimeFilterBetween struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type TimeFilterInput struct {
	And     []*time.Time       `json:"and,omitempty"`
	Or      []*time.Time       `json:"or,omitempty"`
	Not     *TimeFilterInput   `json:"not,omitempty"`
	Eq      *time.Time         `json:"eq,omitempty"`
	Ne      *time.Time         `json:"ne,omitempty"`
	Gt      *time.Time         `json:"gt,omitempty"`
	Gte     *time.Time         `json:"gte,omitempty"`
	Lt      *time.Time         `json:"lt,omitempty"`
	Lte     *time.Time         `json:"lte,omitempty"`
	Null    *bool              `json:"null,omitempty"`
	NotNull *bool              `json:"notNull,omitempty"`
	In      []*time.Time       `json:"in,omitempty"`
	NotIn   []*time.Time       `json:"notIn,omitempty"`
	Between *TimeFilterBetween `json:"between,omitempty"`
}

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

type TodoFiltersInput struct {
	ID        *IDFilterInput      `json:"id,omitempty"`
	Name      *StringFilterInput  `json:"name,omitempty"`
	Users     *UserFiltersInput   `json:"users,omitempty"`
	Owner     *UserFiltersInput   `json:"owner,omitempty"`
	OwnerID   *IDFilterInput      `json:"ownerID,omitempty"`
	CreatedAt *TimeFilterInput    `json:"createdAt,omitempty"`
	UpdatedAt *TimeFilterInput    `json:"updatedAt,omitempty"`
	DeletedAt *TimeFilterInput    `json:"deletedAt,omitempty"`
	And       []*TodoFiltersInput `json:"and,omitempty"`
	Or        []*TodoFiltersInput `json:"or,omitempty"`
	Not       *TodoFiltersInput   `json:"not,omitempty"`
}

type TodoInput struct {
	Name    string    `json:"name"`
	Etype1  *TodoType `json:"etype1,omitempty"`
	Etype5  TodoType  `json:"etype5"`
	Test123 *string   `json:"test123,omitempty"`
}

type TodoOrder struct {
	Asc  *TodoOrderable `json:"asc,omitempty"`
	Desc *TodoOrderable `json:"desc,omitempty"`
}

type TodoPatch struct {
	Name    *string   `json:"name,omitempty"`
	Etype1  *TodoType `json:"etype1,omitempty"`
	Etype5  *TodoType `json:"etype5,omitempty"`
	Test123 *string   `json:"test123,omitempty"`
}

type TodoQueryResult struct {
	Data       []*Todo `json:"data"`
	Count      int     `json:"count"`
	TotalCount int     `json:"totalCount"`
}

type UpdateCatInput struct {
	Filter *CatFiltersInput `json:"filter"`
	Set    *CatPatch        `json:"set"`
}

type UpdateCatPayload struct {
	Cat   *CatQueryResult `json:"cat"`
	Count int             `json:"count"`
}

type UpdateCompanyInput struct {
	Filter *CompanyFiltersInput `json:"filter"`
	Set    *CompanyPatch        `json:"set"`
}

type UpdateCompanyPayload struct {
	Company *CompanyQueryResult `json:"company"`
	Count   int                 `json:"count"`
}

type UpdateSmartPhoneInput struct {
	Filter *SmartPhoneFiltersInput `json:"filter"`
	Set    *SmartPhonePatch        `json:"set"`
}

type UpdateSmartPhonePayload struct {
	SmartPhone *SmartPhoneQueryResult `json:"smartPhone"`
	Count      int                    `json:"count"`
}

type UpdateTodoInput struct {
	Filter *TodoFiltersInput `json:"filter"`
	Set    *TodoPatch        `json:"set"`
}

type UpdateTodoPayload struct {
	Todo  *TodoQueryResult `json:"todo"`
	Count int              `json:"count"`
}

type UpdateUserInput struct {
	Filter *UserFiltersInput `json:"filter"`
	Set    *UserPatch        `json:"set"`
}

type UpdateUserPayload struct {
	User  *UserQueryResult `json:"user"`
	Count int              `json:"count"`
}

type User struct {
	ID          int           `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name        string        `json:"name"`
	CreatedAt   *time.Time    `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time    `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time    `json:"deletedAt,omitempty"`
	Cat         *Cat          `json:"cat,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;;"`
	CompanyID   *int          `json:"companyID,omitempty"`
	Company     *Company      `json:"company,omitempty"`
	SmartPhones []*SmartPhone `json:"smartPhones,omitempty"`
}

type UserFiltersInput struct {
	ID          *IDFilterInput          `json:"id,omitempty"`
	Name        *StringFilterInput      `json:"name,omitempty"`
	CreatedAt   *TimeFilterInput        `json:"createdAt,omitempty"`
	UpdatedAt   *TimeFilterInput        `json:"updatedAt,omitempty"`
	DeletedAt   *TimeFilterInput        `json:"deletedAt,omitempty"`
	Cat         *CatFiltersInput        `json:"cat,omitempty"`
	CompanyID   *IntFilterInput         `json:"companyID,omitempty"`
	Company     *CompanyFiltersInput    `json:"company,omitempty"`
	SmartPhones *SmartPhoneFiltersInput `json:"smartPhones,omitempty"`
	And         []*UserFiltersInput     `json:"and,omitempty"`
	Or          []*UserFiltersInput     `json:"or,omitempty"`
	Not         *UserFiltersInput       `json:"not,omitempty"`
}

type UserInput struct {
	Name        string             `json:"name"`
	Cat         *CatInput          `json:"cat,omitempty"`
	CompanyID   *int               `json:"companyID,omitempty"`
	Company     *CompanyInput      `json:"company,omitempty"`
	SmartPhones []*SmartPhoneInput `json:"smartPhones,omitempty"`
}

type UserOrder struct {
	Asc  *UserOrderable `json:"asc,omitempty"`
	Desc *UserOrderable `json:"desc,omitempty"`
}

type UserPatch struct {
	Name        *string            `json:"name,omitempty"`
	Cat         *CatPatch          `json:"cat,omitempty"`
	CompanyID   *int               `json:"companyID,omitempty"`
	Company     *CompanyPatch      `json:"company,omitempty"`
	SmartPhones []*SmartPhonePatch `json:"smartPhones,omitempty"`
}

type UserQueryResult struct {
	Data       []*User `json:"data"`
	Count      int     `json:"count"`
	TotalCount int     `json:"totalCount"`
}

type UserRef2TodosInput struct {
	Filter *TodoFiltersInput `json:"filter"`
	Set    []int             `json:"set"`
}

type CatOrderable string

const (
	CatOrderableID     CatOrderable = "id"
	CatOrderableName   CatOrderable = "name"
	CatOrderableUserID CatOrderable = "userID"
	CatOrderableAlive  CatOrderable = "alive"
)

var AllCatOrderable = []CatOrderable{
	CatOrderableID,
	CatOrderableName,
	CatOrderableUserID,
	CatOrderableAlive,
}

func (e CatOrderable) IsValid() bool {
	switch e {
	case CatOrderableID, CatOrderableName, CatOrderableUserID, CatOrderableAlive:
		return true
	}
	return false
}

func (e CatOrderable) String() string {
	return string(e)
}

func (e *CatOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CatOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CatOrderable", str)
	}
	return nil
}

func (e CatOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CompanyOrderable string

const (
	CompanyOrderableID              CompanyOrderable = "id"
	CompanyOrderableName            CompanyOrderable = "name"
	CompanyOrderableDescription     CompanyOrderable = "description"
	CompanyOrderableMotherCompanyID CompanyOrderable = "motherCompanyID"
)

var AllCompanyOrderable = []CompanyOrderable{
	CompanyOrderableID,
	CompanyOrderableName,
	CompanyOrderableDescription,
	CompanyOrderableMotherCompanyID,
}

func (e CompanyOrderable) IsValid() bool {
	switch e {
	case CompanyOrderableID, CompanyOrderableName, CompanyOrderableDescription, CompanyOrderableMotherCompanyID:
		return true
	}
	return false
}

func (e CompanyOrderable) String() string {
	return string(e)
}

func (e *CompanyOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CompanyOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CompanyOrderable", str)
	}
	return nil
}

func (e CompanyOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SmartPhoneOrderable string

const (
	SmartPhoneOrderableID          SmartPhoneOrderable = "id"
	SmartPhoneOrderableBrand       SmartPhoneOrderable = "brand"
	SmartPhoneOrderablePhonenumber SmartPhoneOrderable = "phonenumber"
	SmartPhoneOrderableUserID      SmartPhoneOrderable = "userID"
)

var AllSmartPhoneOrderable = []SmartPhoneOrderable{
	SmartPhoneOrderableID,
	SmartPhoneOrderableBrand,
	SmartPhoneOrderablePhonenumber,
	SmartPhoneOrderableUserID,
}

func (e SmartPhoneOrderable) IsValid() bool {
	switch e {
	case SmartPhoneOrderableID, SmartPhoneOrderableBrand, SmartPhoneOrderablePhonenumber, SmartPhoneOrderableUserID:
		return true
	}
	return false
}

func (e SmartPhoneOrderable) String() string {
	return string(e)
}

func (e *SmartPhoneOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SmartPhoneOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SmartPhoneOrderable", str)
	}
	return nil
}

func (e SmartPhoneOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TodoOrderable string

const (
	TodoOrderableID      TodoOrderable = "id"
	TodoOrderableName    TodoOrderable = "name"
	TodoOrderableOwnerID TodoOrderable = "ownerID"
)

var AllTodoOrderable = []TodoOrderable{
	TodoOrderableID,
	TodoOrderableName,
	TodoOrderableOwnerID,
}

func (e TodoOrderable) IsValid() bool {
	switch e {
	case TodoOrderableID, TodoOrderableName, TodoOrderableOwnerID:
		return true
	}
	return false
}

func (e TodoOrderable) String() string {
	return string(e)
}

func (e *TodoOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoOrderable", str)
	}
	return nil
}

func (e TodoOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TodoType string

const (
	TodoTypeBug     TodoType = "Bug"
	TodoTypeFeature TodoType = "Feature"
)

var AllTodoType = []TodoType{
	TodoTypeBug,
	TodoTypeFeature,
}

func (e TodoType) IsValid() bool {
	switch e {
	case TodoTypeBug, TodoTypeFeature:
		return true
	}
	return false
}

func (e TodoType) String() string {
	return string(e)
}

func (e *TodoType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoType", str)
	}
	return nil
}

func (e TodoType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserOrderable string

const (
	UserOrderableID        UserOrderable = "id"
	UserOrderableName      UserOrderable = "name"
	UserOrderableCompanyID UserOrderable = "companyID"
)

var AllUserOrderable = []UserOrderable{
	UserOrderableID,
	UserOrderableName,
	UserOrderableCompanyID,
}

func (e UserOrderable) IsValid() bool {
	switch e {
	case UserOrderableID, UserOrderableName, UserOrderableCompanyID:
		return true
	}
	return false
}

func (e UserOrderable) String() string {
	return string(e)
}

func (e *UserOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserOrderable", str)
	}
	return nil
}

func (e UserOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
