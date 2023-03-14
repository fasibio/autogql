
{{ reserveImport "gorm.io/gorm" }}
{{ reserveImport "context" }}
{{- range $import := .Imports }}
	{{ reserveImport $import }}
{{- end}}
{{- $root := .}}
{{$hookBaseName := "AutoGql"}}

type {{$hookBaseName}}HookM interface {
	{{$root.HookList "model." "" | join "|"}}
}
type {{$hookBaseName}}HookF interface {
	{{$root.HookList "model." "FiltersInput" | join "|"}}
}

type {{$hookBaseName}}HookQueryO interface {
	{{$root.HookList "model." "Order" | join "|"}}
}

type {{$hookBaseName}}HookI interface {
	{{$root.HookList "model." "Input" | join "|"}}
}

type {{$hookBaseName}}HookU interface {
  {{$root.HookList "model.Update" "Input" | join "|"}}
}

type {{$hookBaseName}}HookUP interface {
  {{$root.HookList "model.Update" "Payload" | join "|"}}
}

type {{$hookBaseName}}HookDP interface {
  {{$root.HookList "model.Delete" "Payload" | join "|"}}
}


type {{$hookBaseName}}HookAP interface {
{{$root.HookList "model.Add" "Payload" | join "|"}}
}


type {{$hookBaseName}}DB struct {
	Db *gorm.DB
	Hooks map[string]any
}
func New{{$hookBaseName}}DB(db *gorm.DB) {{$hookBaseName}}DB {
	return {{$hookBaseName}}DB{
		Db:    db,
		Hooks: make(map[string]any),
	}
}

func (db *{{$hookBaseName}}DB) Init() {
	db.Db.AutoMigrate({{.ModelsMigrations}})
}

func AddGetHook[T {{$hookBaseName}}HookM, I any](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookGet[T,I]) {
	db.Hooks[name] = implementation
}

func AddQueryHook[M {{$hookBaseName}}HookM, F {{$hookBaseName}}HookF, O {{$hookBaseName}}HookQueryO](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookQuery[M, F, O]) {
	db.Hooks[name] = implementation
}

func AddAddHook[M {{$hookBaseName}}HookM,I {{$hookBaseName}}HookI, AP {{$hookBaseName}}HookAP](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookAdd[M, I, AP]) {
	db.Hooks[name] = implementation
}

func AddUpdateHook[M {{$hookBaseName}}HookM, U {{$hookBaseName}}HookU, UP {{$hookBaseName}}HookUP](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookUpdate[U, UP]) {
	db.Hooks[name] = implementation
}

func AddDeleteHook[F {{$hookBaseName}}HookF, DP {{$hookBaseName}}HookDP](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookDelete[F, DP]) {
	db.Hooks[name] = implementation
}

type {{$hookBaseName}}HookGet[obj {{$hookBaseName}}HookM, identifier any] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, id ...identifier) (*gorm.DB, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data *obj) (*obj, error)
	BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error)
}

type DefaultGetHook[obj {{$hookBaseName}}HookM, identifier any] struct{}

func (d DefaultGetHook[obj, identifier]) Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, id ...identifier) (*gorm.DB, error) {
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

type {{$hookBaseName}}HookQuery[obj {{$hookBaseName}}HookM, filter {{$hookBaseName}}HookF, order {{$hookBaseName}}HookQueryO] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, filter *filter, order *order, first, offset *int) (*gorm.DB, *filter, *order, *int, *int, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error)
	BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error)
}

type DefaultQueryHook[obj {{$hookBaseName}}HookM, filter {{$hookBaseName}}HookF, order {{$hookBaseName}}HookQueryO] struct{}

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

type {{$hookBaseName}}HookAdd[obj {{$hookBaseName}}HookM, input {{$hookBaseName}}HookI, res {{$hookBaseName}}HookAP] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input []*input) (*gorm.DB, []*input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB,[]obj, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type DefaultAddHook[obj {{$hookBaseName}}HookM, input {{$hookBaseName}}HookI, res {{$hookBaseName}}HookAP] struct{}

func (d DefaultAddHook[obj, inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input []*inputType) (*gorm.DB, []*inputType, error) {
	return dbHelper.Db, input, nil
}
func (d DefaultAddHook[obj, inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB, []obj, error) {
	return db, data, nil
}
func (d DefaultAddHook[obj, inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}

type {{$hookBaseName}}HookUpdate[ input {{$hookBaseName}}HookU,  res {{$hookBaseName}}HookUP]interface{
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type DefaultUpdateHook[input {{$hookBaseName}}HookU, res {{$hookBaseName}}HookUP] struct{}

func (d DefaultUpdateHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}
func (d DefaultUpdateHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error) {
	return db, data, nil
}
func (d DefaultUpdateHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}

type {{$hookBaseName}}HookDelete[input {{$hookBaseName}}HookF, res {{$hookBaseName}}HookDP] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type DefaultDeleteHook[input {{$hookBaseName}}HookF, res {{$hookBaseName}}HookDP] struct{}

func (d DefaultDeleteHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}
func (d DefaultDeleteHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}
func (d DefaultDeleteHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}
