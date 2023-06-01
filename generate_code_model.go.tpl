
{{ reserveImport "time"  }}
{{- $root := .}}
{{- $input2TypeName := "MergeToType"}}
{{- range $objectName, $object := .Handler.List.Enums }}
  func (d *{{$objectName}}) {{$input2TypeName}}() {{$objectName}} {
    return *d
  }
{{- end}}

{{- range $objectName, $object := .Handler.List.Objects }}
  {{$objectName := $object.Name}}

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
