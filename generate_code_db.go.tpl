
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

func AddGetHook[T {{$hookBaseName}}HookM](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookGet[T]) {
	db.Hooks[name] = implementation
}

func AddQueryHook[M {{$hookBaseName}}HookM, F {{$hookBaseName}}HookF, O {{$hookBaseName}}HookQueryO](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookQuery[M, F, O]) {
	db.Hooks[name] = implementation
}

func AddAddHook[M {{$hookBaseName}}HookM,I {{$hookBaseName}}HookI, AP {{$hookBaseName}}HookAP](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookAdd[M, I, AP]) {
	db.Hooks[name] = implementation
}

func AddUpdateHook[M {{$hookBaseName}}HookM, U {{$hookBaseName}}HookU, UP {{$hookBaseName}}HookUP](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookUpdate[M, U, UP]) {
	db.Hooks[name] = implementation
}

func AddDeleteHook[M {{$hookBaseName}}HookM, F {{$hookBaseName}}HookF, DP {{$hookBaseName}}HookDP](db *{{$hookBaseName}}DB, name string, implementation {{$hookBaseName}}HookDelete[M, F, DP]) {
	db.Hooks[name] = implementation
}

type {{$hookBaseName}}HookGet[obj {{$hookBaseName}}HookM] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, id int) (*gorm.DB, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data *obj) (*obj, error)
	BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error)
}

type {{$hookBaseName}}HookQuery[obj {{$hookBaseName}}HookM, filter {{$hookBaseName}}HookF, order {{$hookBaseName}}HookQueryO] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, filter *filter, order *order, first, offset *int) (*gorm.DB, *filter, *order, *int, *int, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error)
	BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error)
}

type {{$hookBaseName}}HookAdd[obj {{$hookBaseName}}HookM, input {{$hookBaseName}}HookI, res {{$hookBaseName}}HookAP] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input []*input) (*gorm.DB, []*input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB,[]obj, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type {{$hookBaseName}}HookUpdate[obj {{$hookBaseName}}HookM, input {{$hookBaseName}}HookU,  res {{$hookBaseName}}HookUP]interface{
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB, data *obj) (*gorm.DB, *obj, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}

type {{$hookBaseName}}HookDelete[obj {{$hookBaseName}}HookM, input {{$hookBaseName}}HookF, res {{$hookBaseName}}HookDP] interface {
	Received(ctx context.Context, dbHelper *{{$hookBaseName}}DB, input *input) (*gorm.DB, input, error)
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)
}