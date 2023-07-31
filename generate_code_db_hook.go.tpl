
{{ reserveImport "gorm.io/gorm" }}
{{ reserveImport "context" }}
{{- range $import := .Imports }}
  {{ reserveImport $import }}
{{- end}}
{{- $root := .}}
{{$hookBaseName := "AutoGql"}}
type QueryName string
type GetName string
type AddName string
type UpdateName string
type DeleteName string
type Many2ManyName string 

const (
  {{- range $objectName, $object := .Handler.List.Objects }}
    {{- if $object.HasSqlDirective}}
      {{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
        Add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s Many2ManyName = "Add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s"
      {{- end}}
      {{- if $object.SQLDirective.Query.Get}}
        Get{{$object.Name}} GetName = "Get{{$object.Name}}"
      {{- end}}
      {{- if $object.SQLDirective.Query.Query}}
        Query{{$object.Name}} QueryName = "Query{{$object.Name}}"
      {{- end}}
      {{- if $object.SQLDirective.Mutation.Add}}
        Add{{$object.Name}} AddName = "Add{{$object.Name}}"
      {{- end}}
      {{- if $object.SQLDirective.Mutation.Update}}
        Update{{$object.Name}} UpdateName = "Update{{$object.Name}}"
      {{- end}}
      {{- if $object.SQLDirective.Mutation.Delete}}
        Delete{{$object.Name}} DeleteName = "Delete{{$object.Name}}"
      {{- end}}
    {{- end}}
  {{- end}}
)

// Modelhooks
type {{$hookBaseName}}HookM interface {
  {{$root.HookList "model." "" | join "|"}}
}

// Filter Hooks
type {{$hookBaseName}}HookF interface {
  {{$root.HookList "model." "FiltersInput" | join "|"}}
}


{{- $m2mV := $root.HookListMany2Many "model."}}
// Many2Many Hooks
type {{$hookBaseName}}HookM2M interface {
  {{$m2mV | join "|"}}
}

// Order Hooks
type {{$hookBaseName}}HookQueryO interface {
  {{$root.HookList "model." "Order" | join "|"}}
}

// Input Hooks
type {{$hookBaseName}}HookI interface {
  {{$root.HookList "model." "Input" | join "|"}}
}

// Update Hooks
type {{$hookBaseName}}HookU interface {
  {{$root.HookList "model.Update" "Input" | join "|"}}
}

// Update Payload Hooks
type {{$hookBaseName}}HookUP interface {
  {{$root.HookList "model.Update" "Payload" | join "|"}}
}

// Delete Payload Hooks
type {{$hookBaseName}}HookDP interface {
  {{$root.HookList "model.Delete" "Payload" | join "|"}}
}

// Add Payload Hooks
type {{$hookBaseName}}HookAP interface {
{{$root.HookList "model.Add" "Payload" | join "|"}}
}

// Add a getHook
// useable for 
{{- range $objectName, $object := .Handler.List.Objects }}  
  {{- if $object.HasSqlDirective}} 
  {{- if $object.SQLDirective.Query.Get}}
//  - Get{{$object.Name}}
  {{- end}} 
  {{- end}} 
{{- end }}
func AddGetHook[T {{$hookBaseName}}HookM, I any](db *{{$hookBaseName}}DB, name GetName, implementation {{$hookBaseName}}HookGet[T,I]) {
  db.Hooks[string(name)] = implementation
}

// Add a queryHook
// useable for 
{{- range $objectName, $object := .Handler.List.Objects }}  
  {{- if $object.HasSqlDirective}} 
  {{- if $object.SQLDirective.Query.Query}}
//  - Query{{$object.Name}}
  {{- end}} 
  {{- end}} 
{{- end }}
func AddQueryHook[M {{$hookBaseName}}HookM, F {{$hookBaseName}}HookF, O {{$hookBaseName}}HookQueryO](db *{{$hookBaseName}}DB, name QueryName, implementation {{$hookBaseName}}HookQuery[M, F, O]) {
  db.Hooks[string(name)] = implementation
}

// Add a addHook
// useable for 
{{- range $objectName, $object := .Handler.List.Objects }}  
  {{- if $object.HasSqlDirective}} 
  {{- if $object.SQLDirective.Mutation.Add}}
//  - Add{{$object.Name}}
  {{- end}} 
  {{- end}} 
{{- end }}
func AddAddHook[M {{$hookBaseName}}HookM,I {{$hookBaseName}}HookI, AP {{$hookBaseName}}HookAP](db *{{$hookBaseName}}DB, name AddName, implementation {{$hookBaseName}}HookAdd[M, I, AP]) {
  db.Hooks[string(name)] = implementation
}

// Add a updateHook
// useable for 
{{- range $objectName, $object := .Handler.List.Objects }}  
  {{- if $object.HasSqlDirective}} 
  {{- if $object.SQLDirective.Mutation.Update}}
//  - Update{{$object.Name}}
  {{- end}} 
  {{- end}} 
{{- end }}
func AddUpdateHook[M {{$hookBaseName}}HookM, U {{$hookBaseName}}HookU, UP {{$hookBaseName}}HookUP](db *{{$hookBaseName}}DB, name UpdateName, implementation {{$hookBaseName}}HookUpdate[U, UP]) {
  db.Hooks[string(name)] = implementation
}

// Add a Many2Many Hook
// useable for 
{{- range $objectName, $object := .Handler.List.Objects }}  
  {{- if $object.HasSqlDirective}} 
  {{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
//  - Add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s
  {{- end}} 
  {{- end}} 
{{- end }}
func AddMany2ManyHook[U AutoGqlHookM2M, UP AutoGqlHookUP](db *AutoGqlDB, name Many2ManyName, implementation AutoGqlHookMany2Many[U, UP]) {
  db.Hooks[string(name)] = implementation
}

// Add a updateHook
// useable for 
{{- range $objectName, $object := .Handler.List.Objects }}  
  {{- if $object.HasSqlDirective}} 
  {{- if $object.SQLDirective.Mutation.Delete}}
//  - Delete{{$object.Name}}
  {{- end}} 
  {{- end}} 
{{- end }}
func AddDeleteHook[F {{$hookBaseName}}HookF, DP {{$hookBaseName}}HookDP](db *{{$hookBaseName}}DB, name DeleteName, implementation {{$hookBaseName}}HookDelete[F, DP]) {
  db.Hooks[string(name)] = implementation
}

// Interface description of a getHook
// Simple you can use DefaultGetHook and only implement the hooks you need: 
//   type MyGetHook struct {
//      DefaultGetHook[model.Todo, model.TodoInput, model.AddTodoPayload]
//   }
//   func (m MyGetHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Todo) (*gorm.DB, []model.Todo, error) {
//      //do some stuff
//      return db, data, nil
//   }
type {{$hookBaseName}}HookGet[obj {{$hookBaseName}}HookM, identifier any] interface {
  Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, id ...identifier) (*gorm.DB, error) // Direct after Resolver is call
  BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) // Direct before call Database
  AfterCallDb(ctx context.Context, data *obj) (*obj, error) // After database call with resultset from database
  BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error) // short before return the data
}

// Default get hook implementation
// Simple you can use and only implement the hooks you need: 
//   type MyGetHook struct {
//      DefaultGetHook[model.Todo, int64]
//   }
//   func (m MyGetHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, []model.Todo, error) {
//      //do some stuff
//      return db, data, nil
//   }
type DefaultGetHook[obj {{$hookBaseName}}HookM, identifier any] struct{}

// Direct after Resolver is call
func (d DefaultGetHook[obj, identifier]) Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, id ...identifier) (*gorm.DB, error) {
  return dbHelper.Db, nil
}

// Direct before call Database
func (d DefaultGetHook[obj, identifier]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
  return db, nil
}

// After database call with resultset from database
func (d DefaultGetHook[obj, identifier]) AfterCallDb(ctx context.Context, data *obj) (*obj, error) {
  return data, nil
}

 // short before return the data
func (d DefaultGetHook[obj, identifier]) BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error) {
  return data, nil
}

// Interface description of a query Hook
// Simple you can use DefaultQueryHook and only implement the hooks you need: 
//   type MyQueryHook struct {
//      DefaultQueryHook[model.Todo, model.TodoFiltersInput, model.TodoOrder]
//   }
//   func (m MyQueryHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, []model.Todo, error) {
//      //do some stuff
//      return db, nil
//   }
type {{$hookBaseName}}HookQuery[obj {{$hookBaseName}}HookM, filter {{$hookBaseName}}HookF, order {{$hookBaseName}}HookQueryO] interface {
  Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, filter *filter, order *order, first, offset *int) (*gorm.DB, *filter, *order, *int, *int, error) // Direct after Resolver is call
  BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) // Direct before call Database
  AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error) // After database call with resultset from database
  BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error)  // short before return the data
}

// Default query hook implementation
// Simple you can use DefaultQueryHook and only implement the hooks you need: 
//   type MyQueryHook struct {
//      DefaultQueryHook[model.Todo, model.TodoFiltersInput, model.TodoOrder]
//   }
//   func (m MyQueryHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, []model.Todo, error) {
//      //do some stuff
//      return db, nil
//   }
type DefaultQueryHook[obj {{$hookBaseName}}HookM, filter {{$hookBaseName}}HookF, order {{$hookBaseName}}HookQueryO] struct{}

// Direct after Resolver is call
func (d DefaultQueryHook[obj, filterType, orderType]) Received(ctx context.Context, dbHelper *AutoGqlDB, filter *filterType, order *orderType, first, offset *int) (*gorm.DB, *filterType, *orderType, *int, *int, error) {
  return dbHelper.Db, filter, order, first, offset, nil
}

// Direct before call Database
func (d DefaultQueryHook[obj, filter, order]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
  return db, nil
}

// After database call with resultset from database
func (d DefaultQueryHook[obj, filter, order]) AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error) {
  return data, nil
}

 // short before return the data
func (d DefaultQueryHook[obj, filter, order]) BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error) {
  return data, nil
}

// Interface description of a add Hook
// Simple you can use DefaultAddHook and only implement the hooks you need: 
//   type MyAddHook struct {
//      DefaultAddHook[model.Todo, model.TodoInput, model.AddTodoPayload]
//   }
//   func (m MyAddHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Todo) (*gorm.DB, []model.Todo, error) {
//      //do some stuff
//      return db, data, nil
//   }
type {{$hookBaseName}}HookAdd[obj {{$hookBaseName}}HookM, input {{$hookBaseName}}HookI, res {{$hookBaseName}}HookAP] interface {
  Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input []*input) (*gorm.DB, []*input, error) // Direct after Resolver is call
  BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB,[]obj, error) // Direct before call Database
  BeforeReturn(ctx context.Context, db *gorm.DB, data []obj, res *res) (*res, error) // After database call with resultset from database
}

// Default add hook implementation
// Simple you can use DefaultAddHook and only implement the hooks you need: 
//   type MyAddHook struct {
//      DefaultAddHook[model.Todo, model.TodoInput, model.AddTodoPayload]
//   }
//   func (m MyAddHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Todo) (*gorm.DB, []model.Todo, error) {
//      //do some stuff
//      return db, data, nil
//   }
type DefaultAddHook[obj {{$hookBaseName}}HookM, input {{$hookBaseName}}HookI, res {{$hookBaseName}}HookAP] struct{}

// Direct after Resolver is call
func (d DefaultAddHook[obj, inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input []*inputType) (*gorm.DB, []*inputType, error) {
  return dbHelper.Db, input, nil
}

// Direct before call Database
func (d DefaultAddHook[obj, inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB, []obj, error) {
  return db, data, nil
}

// After database call with resultset from database
func (d DefaultAddHook[obj, inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, data []obj, res *resType) (*resType, error) {
  return res, nil
}

// Interface description of a update Hook
// Simple you can use DefaultUpdateHook and only implement the hooks you need: 
//   type MyUpdateHook struct {
//      DefaultUpdateHook[model.TodoInput, model.UpdateTodoPayload]
//   }
//   func (m MyUpdateHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB,  map[string]interface{}, error) {
//      //do some stuff
//      return db, data, nil
//   }
type {{$hookBaseName}}HookUpdate[ input {{$hookBaseName}}HookU,  res {{$hookBaseName}}HookUP]interface{
  Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input *input) (*gorm.DB, input, error) // Direct after Resolver is call
  BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error) // Direct before call Database
  BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error) // After database call with resultset from database
}

// Default update hook implementation
// Simple you can use DefaultUpdateHook and only implement the hooks you need: 
//   type MyUpdateHook struct {
//      DefaultUpdateHook[model.TodoInput, model.UpdateTodoPayload]
//   }
//   func (m MyUpdateHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB,  map[string]interface{}, error) {
//      //do some stuff
//      return db, data, nil
//   }
type DefaultUpdateHook[input {{$hookBaseName}}HookU, res {{$hookBaseName}}HookUP] struct{}

// Direct after Resolver is call
func (d DefaultUpdateHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
  return dbHelper.Db, *input, nil
}

// Direct before call Database
func (d DefaultUpdateHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error) {
  return db, data, nil
}

// After database call with resultset from database
func (d DefaultUpdateHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
  return res, nil
}

// Interface description of a many2many Hook
// Simple you can use DefaultMany2ManyHook and only implement the hooks you need: 
//   type MyM2mHook struct {
//      DefaultMany2ManyHook[model.UserRef2TodosInput, model.UpdateTodoPayload]
//   }
//   func (m MyM2mHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
//      //do some stuff
//      return db, nil
//   }
type {{$hookBaseName}}HookMany2Many[input {{$hookBaseName}}HookM2M, res AutoGqlHookUP] interface {
  Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error) // Direct after Resolver is call
  BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) // Direct before call Database
  BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error) // After database call with resultset from database
}

// Default many2many hook implementation
// Simple you can use DefaultMany2ManyHook and only implement the hooks you need: 
//   type MyM2mHook struct {
//      DefaultMany2ManyHook[model.UserRef2TodosInput, model.UpdateTodoPayload]
//   }
//   func (m MyM2mHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
//      //do some stuff
//      return db, nil
//   }
type DefaultMany2ManyHook[input {{$hookBaseName}}HookM2M, res {{$hookBaseName}}HookUP] struct {}

// Direct after Resolver is call
func (d DefaultMany2ManyHook[inputType, resType])Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input *inputType) (*gorm.DB, inputType, error){
  return dbHelper.Db, *input, nil
}

// Direct before call Database
func (d DefaultMany2ManyHook[inputType, resType])BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error){
  return db, nil
}

// After database call with resultset from database
func (d DefaultMany2ManyHook[inputType, resType])BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error){
  return res, nil
}

// Interface description of a delete Hook
// Simple you can use DefaultDeleteHook and only implement the hooks you need: 
//   type MyM2mHook struct {
//      DefaultDeleteHook[model.TodoFiltersInput, model.DeleteTodoPayload]
//   }
//   func (m MyM2mHook) BeforeCallDb(ctx context.Context, db *gorm.DB, input model.TodoFiltersInput) (*gorm.DB, model.TodoFiltersInput, error) {
//      //do some stuff
//      return db, input, nil
//   }
type {{$hookBaseName}}HookDelete[input {{$hookBaseName}}HookF, res {{$hookBaseName}}HookDP] interface {
  Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input *input) (*gorm.DB, input, error) // Direct after Resolver is call
  BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) // Direct before call Database
  BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error) // After database call with resultset from database
}

// Default delete hook implementation
// Simple you can use DefaultDeleteHook and only implement the hooks you need: 
//   type MyM2mHook struct {
//      DefaultDeleteHook[model.TodoFiltersInput, model.DeleteTodoPayload]
//   }
//   func (m MyM2mHook) BeforeCallDb(ctx context.Context, db *gorm.DB, input model.TodoFiltersInput) (*gorm.DB, model.TodoFiltersInput, error) {
//      //do some stuff
//      return db, input, nil
//   }
type DefaultDeleteHook[input {{$hookBaseName}}HookF, res {{$hookBaseName}}HookDP] struct{}

// Direct after Resolver is call
func (d DefaultDeleteHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
  return dbHelper.Db, *input, nil
}

// Direct before call Database
func (d DefaultDeleteHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
  return db, nil
}

// After database call with resultset from database
func (d DefaultDeleteHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
  return res, nil
}
