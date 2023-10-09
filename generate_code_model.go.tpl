
{{ reserveImport "fmt"  }}
{{ reserveImport "time"  }}
{{ reserveImport "io"  }}
{{ reserveImport "github.com/99designs/gqlgen/graphql"  }}
{{ reserveImport "github.com/mitchellh/mapstructure"  }}
{{ reserveImport "gorm.io/gorm"  }}

{{- $root := .}}
{{- $input2TypeName := "MergeToType"}}

// GetInputStruct returns struct filled from map obj defined by name
// Example useage struct validation with github.com/go-playground/validator by directive: 
//    func ValidateDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
//      field := graphql.GetPathContext(ctx)
//      if data, ok := obj.(map[string]interface{}); ok {
//        for _, v := range field.ParentField.Field.Arguments {
//          name := v.Value.ExpectedType.Name()
//          model, err := model.GetInputStruct(name, data)
//          if err != nil {
//            //handle not found error
//          }
//          if err := validate.Struct(model); err != nil {
//          //handle error
//          }
//        }
//      }
//      return next(ctx)
//    }
func GetInputStruct(name string, obj map[string]interface{}) (interface{}, error) {
	switch name {
    {{- range $objectName, $object := .Handler.List.Objects }}
      {{- if $object.HasSqlDirective}}
        case "{{$objectName}}Input":
          return {{$objectName}}InputFromMap(obj)
      {{- end}}
    {{- end}}
  }
	return nil, fmt.Errorf("%s not found", name)
}

{{- range $objectName, $object := .Handler.List.Enums }}
  // {{$input2TypeName}} for enum value {{$objectName}}
  func (d *{{$objectName}}) {{$input2TypeName}}() {{$objectName}} {
    return *d
  }
{{- end}}

{{- range $objectName, $object := .Handler.List.Objects }}
  {{- if $object.HasSqlDirective}}
    {{$objectName := $object.Name}}
    // {{$objectName}}InputFromMap return a {{$objectName}}Input from data map
    // use github.com/mitchellh/mapstructure with reflaction
    func {{$objectName}}InputFromMap(data map[string]interface{}) ({{$objectName}}Input, error) {
        model := {{$objectName}}Input{}
        err := mapstructure.Decode(data, &model); 
        return model, err
    }

    // {{$input2TypeName}} returns a map with all values set to {{$objectName}}Patch
    func (d *{{$objectName}}Patch) {{$input2TypeName}}() map[string]interface{} {
      res := make(map[string]interface{})

      {{- range $entityKey, $entity := $object.PatchEntities }}
        {{- $entityGoName :=  $root.GetGoFieldName $objectName $entity}}
        {{- if  $entity.IsPrimitive}} 
          if d.{{$entityGoName}} != nil {
            {{- if $entity.IsArray}}
              var tmp{{$entityGoName}} {{$root.GetGoFieldType $objectName $entity false}} 
              for _, v := range d.{{$entityGoName}}{
                tmp := v
                tmp{{$entityGoName}} = append(tmp{{$entityGoName}}, {{$root.GenPointerStrIfNeeded $objectName $entity false}}tmp)
              }
              res["{{$entity.DatabaseFieldName}}"] = tmp{{$entityGoName}}
            {{- else}}
              res["{{$entity.DatabaseFieldName}}"] = {{$root.PointerStrIfNeeded $objectName $entity true}}d.{{$entityGoName}}
            {{- end}}	
          }
        {{- else}}
          if d.{{$entityGoName}} != nil {
            {{- if $entity.IsArray}}
              tmp{{$entityGoName}} := make([]map[string]interface{},len(d.{{$entityGoName}}))
              for _, v := range d.{{$entityGoName}}{
                tmp := v.{{$input2TypeName}}()
                tmp{{$entityGoName}} = append(tmp{{$entityGoName}}, tmp)
              }
              res["{{$entity.DatabaseFieldName}}"] = tmp{{$entityGoName}}
            {{- else}}
              {{- if $entity.IsScalar}}
                res["{{$entity.DatabaseFieldName}}"] = {{if not $entity.Required}}*{{end}} d.{{$entityGoName}}
              {{- else}}
                res["{{$entity.DatabaseFieldName}}"] = d.{{$entityGoName}}.{{$input2TypeName}}()
              {{- end}}
            {{- end}}	
        }
        {{- end}}
      {{- end}}
      return res
    }

    // {{$input2TypeName}} retuns a {{$objectName}} filled from {{$objectName}}Input
    func (d *{{$objectName}}Input) {{$input2TypeName}}() {{$objectName}} {
      {{- range $entityKey, $entity := $object.InputEntities }}
        {{- $entityGoName := $root.GetGoFieldName $objectName $entity}}
        {{- if not $entity.IsPrimitive}}
          {{$pointer := "*"}}
          {{- if $entity.Required}}
            {{$pointer := ""}}
          {{- end}}
          {{$entityType := $root.GetGoFieldType $objectName $entity true}}
          var tmp{{$entityGoName}} {{ if $entity.IsArray}} []{{$pointer}}{{$entityType }} {{ else }} {{$entityType }} {{ end }}
          {{- if not $entity.Required}}
            if d.{{$entityGoName}} != nil {
          {{- end}}
          {{- if $entity.IsArray}}
            tmp{{$entityGoName}} = make([]*{{$entityType }},len(d.{{$entityGoName}}))
            for _, v := range d.{{$entityGoName}}{
              tmp := v.{{$input2TypeName}}()
              tmp{{$entityGoName}} = append(tmp{{$entityGoName}}, &tmp)
            }
          {{- else}}
            {{- if $entity.IsScalar}}
              tmp{{$entityGoName}} = {{if not $entity.Required}}*{{end}}d.{{$entityGoName}}
            {{- else}}
              tmp{{$entityGoName}} = d.{{$entityGoName}}.{{$input2TypeName}}()
            {{- end}}
          {{- end}}	
          {{- if not $entity.Required}}
            }
          {{- end}}
        {{- else}}
          {{$entityType := $root.GetGoFieldType $objectName $entity false}}
          {{- if $entity.Required}}
            {{- if and $entity.IsArray (not $entity.IsArrayElementRequired)}}
              tmp{{$entityGoName}} := make({{$entityType }}, len(d.{{$entityGoName}}))			
              for _, v := range d.{{$entityGoName}}{
                tmp{{$entityGoName}} = append(tmp{{$entityGoName}}, &v)
              }
            {{- else}}
              tmp{{$entityGoName}} := d.{{$entityGoName}}
            {{- end}}
          {{- else}}		
            var tmp{{$entityGoName}} {{$entityType }}
            if d.{{$entityGoName}} != nil {
              tmp{{$entityGoName}} = d.{{$entityGoName}}
            }
          {{- end}}

        {{- end}}
      {{- end}}
      return {{$objectName}}{
      {{- range $entityKey, $entity := $object.InputEntities }}
        {{- $entityGoName := $root.GetGoFieldName $objectName $entity}}
        {{$entityGoName}}: {{$root.GetPointerSymbol $objectName $entity}}tmp{{$entityGoName}},
      {{- end}}
      }
    }
  {{- end}}
{{- end}}


func MarshalDeletedAt(d gorm.DeletedAt) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
    w.Write([]byte(d.Time.String()))
	})
}

func UnmarshalDeletedAt(v interface{}) (gorm.DeletedAt, error) {
	switch v := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", v)
    if err != nil {
      return gorm.DeletedAt{},err 
    }
    return gorm.DeletedAt{
      Time: t,
			Valid: true,
    },nil
	default:
		return gorm.DeletedAt{}, fmt.Errorf("%T is not a gorm.DeletedAt", v)
	}
}

func MarshalInputDeletedAt(d gorm.DeletedAt) graphql.Marshaler {
	return MarshalDeletedAt(d)
}

func UnmarshalInputDeletedAt(v interface{}) (gorm.DeletedAt, error) {
  return UnmarshalDeletedAt(v)
}