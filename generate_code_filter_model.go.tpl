{{ reserveImport "fmt"  }}
{{ reserveImport "gorm.io/gorm"  }}
{{ reserveImport "strings" }}
{{ reserveImport "github.com/fasibio/autogql/runtimehelper" }}

	
{{$methodeName := "ExtendsDatabaseQuery"}}


{{- $root := .}}

{{- range $objectName, $object := .Handler.List.Objects }}
{{- if $object.HasSqlDirective}}

func (d *{{$object.Name}}FiltersInput) {{$methodeName}}(db *gorm.DB, alias string) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd,v.ExtendsDatabaseQuery(db, alias)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd,tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			
			tmp  = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr,tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot,d.Not.ExtendsDatabaseQuery(db, alias)...))
	}
  {{- range $entityKey, $entity := $object.Entities }}
  {{- $entityGoName :=  $root.GetGoFieldName $objectName $entity}}

	{{-  if or $entity.BuiltIn $entity.GqlTypeObj.HasSqlDirective }}
	if d.{{$entityGoName}} != nil {
    {{-  if $entity.BuiltIn  }}
    res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, fmt.Sprintf("%s.%s",alias,"{{snakecase $entityGoName}}"))...)
    {{- else }}
			{{- if $entity.HasMany2ManyDirective}}
    tableName := db.Config.NamingStrategy.TableName("{{$root.GetGoFieldTypeName $objectName $entity }}")
			{{- $m2mTableName := $entity.Many2ManyDirectiveTable}}
		db := db.Joins(fmt.Sprintf("JOIN {{$m2mTableName}} ON {{$m2mTableName}}.{{$object.Name | snakecase}}_{{$root.PrimaryKeyOfObject $object.Name}} = %s.{{$root.PrimaryKeyOfObject $object.Name}} JOIN %s ON {{$m2mTableName}}.{{$entity.GqlTypeName | snakecase}}_{{$root.PrimaryKeyOfObject $entity.GqlTypeName | snakecase}} = %s.{{$root.PrimaryKeyOfObject $object.Name}}", alias, tableName,tableName))
    res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, tableName)...)
			{{- else}}
    db := db.Joins("{{$entityGoName}}")
		res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, "{{$entityGoName}}")...)
			{{- end}}
    {{- end}}
	}
	{{- end}}
  {{- end}}

	return res
}
{{- end}}
{{- end}}

func (d *StringFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp= append(tmp, runtimehelper.Equal(fieldName,*v))
		}
		res = append(res,  tmp...)
	}
	if d.Contains != nil {
		res = append(res, runtimehelper.Like(fieldName,fmt.Sprintf("%%%s%%", *d.Contains)))
	}

	if d.Containsi != nil {
		res = append(res, runtimehelper.Like(fmt.Sprintf("lower(%s)",fieldName),fmt.Sprintf("%%%s%%", strings.ToLower(*d.Containsi))))
	}

	if d.EndsWith != nil {
		res = append(res, runtimehelper.Like(fieldName,fmt.Sprintf("%%%s", *d.EndsWith)))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName,*d.Eq))
	}

	if d.Eqi != nil {
		res = append(res, runtimehelper.Equal(fmt.Sprintf("lower(%s)",fieldName),strings.ToLower(*d.Eqi)))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName,d.In))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName,*d.Ne))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot,d.Not.ExtendsDatabaseQuery(db,fieldName)...))
	}

	if d.NotContains != nil {
		res = append(res, runtimehelper.NotLike(fieldName, fmt.Sprintf("%%%s%%", *d.NotContains)))
	}

	if d.NotContainsi != nil {
		res = append(res, runtimehelper.NotLike(fmt.Sprintf("lower(%s)",fieldName), fmt.Sprintf("%%%s%%", strings.ToLower(*d.NotContainsi))))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))
	}

	if d.NotNull != nil {
		res = append(res, runtimehelper.NotNull(fieldName, d.NotNull))
	}

	if d.Null != nil {
		res = append(res, runtimehelper.Null(fieldName, d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp= append(tmp, runtimehelper.Equal(fieldName,*v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr,tmp...))
	}

	if d.StartsWith != nil {
		res = append(res, runtimehelper.Like(fieldName,fmt.Sprintf("%s%%", *d.StartsWith)))
	}

	return res
}

func (d *IntFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp= append(tmp, runtimehelper.Equal(fieldName,*v))
		}
		res = append(res,  tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName,d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName,*d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName,*d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName,*d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName,d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName,*d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName,*d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName,*d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot,d.Not.ExtendsDatabaseQuery(db,fieldName)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName,d.NotIn))

	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName,*d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName,*d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp= append(tmp, runtimehelper.Equal(fieldName,*v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr,tmp...))
	}

	return res
}

func (d *BooleanFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Is != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Is))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName)...))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName, *d.NotNull))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName, *d.Null))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}
