
{{ reserveImport "gorm.io/gorm" }}
{{ reserveImport "context" }}
{{- range $import := .Imports }}
  {{ reserveImport $import }}
{{- end}}
{{$baseName := "AutoGql"}}

type {{$baseName}}DB struct {
  Db *gorm.DB
  Hooks map[string]any
}
func New{{$baseName}}DB(db *gorm.DB) {{$baseName}}DB {
  return {{$baseName}}DB{
    Db:    db,
    Hooks: make(map[string]any),
  }
}

func (db *{{$baseName}}DB) Init() {
  db.Db.AutoMigrate({{.ModelsMigrations}})
}

