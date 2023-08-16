
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

// Create a new {{$baseName}}DB
func New{{$baseName}}DB(db *gorm.DB) {{$baseName}}DB {
  return {{$baseName}}DB{
    Db:    db,
    Hooks: make(map[string]any),
  }
}

//execute Gorm AutoMigrate with all @SQL Graphql Types
func (db *{{$baseName}}DB) Init() error {
  return db.Db.AutoMigrate({{.ModelsMigrations}})
}

