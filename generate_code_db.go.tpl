
{{ reserveImport "gorm.io/gorm" }}
{{ reserveImport "context" }}
{{- range $import := .Imports }}
	{{ reserveImport $import }}
{{- end}}
{{$hookBaseName := "AutoGql"}}

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

