{{ reserveImport "fmt"  }}
{{ reserveImport "gorm.io/gorm"  }}
{{ reserveImport "strings" }}
{{ reserveImport "github.com/fasibio/autogql/runtimehelper" }}
	
{{$methodeName := "ExtendsDatabaseQuery"}}
const extendsDatabaseFieldNameFormat string = "%[2]s.%[1]s%[3]s%[1]s"

{{- $root := .}}

{{- range $objectName, $object := .Handler.List.Objects }}
{{- if $object.HasSqlDirective}}

func (d *{{$object.Name}}FiltersInput) PrimaryKeyName() string {
	return "{{$root.PrimaryKeyOfObject $object.Name}}"
}

func (d *{{$object.Name}}FiltersInput) {{$methodeName}}(db *gorm.DB, alias string,deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd,v.ExtendsDatabaseQuery(db, alias, true,blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd,tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			
			tmp  = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true,blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr,tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot,d.Not.ExtendsDatabaseQuery(db, alias, true,blackList)...))
	}
  {{- range $entityKey, $entity := $object.InputFilterEntities }}
  {{- $entityGoName :=  $root.GetGoFieldName $objectName $entity}}

	{{-  if or $entity.IsPrimitive $entity.GqlTypeObj.HasSqlDirective $entity.IsEnum }}
	if d.{{$entityGoName}} != nil {
    {{-  if or $entity.IsPrimitive $entity.IsEnum }}
    res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, fmt.Sprintf(extendsDatabaseFieldNameFormat,runtimehelper.GetQuoteChar(db), alias,"{{snakecase $entityGoName}}"),true,blackList)...)
    {{- else }}
			{{- if $entity.HasMany2ManyDirective}}
    tableName := db.Config.NamingStrategy.TableName("{{$root.GetGoFieldTypeName $objectName $entity }}")
		{{- $m2mTableName := $entity.Many2ManyDirectiveTable}}
		if _, ok := blackList["{{$m2mTableName}}"]; !ok {
			blackList["{{$m2mTableName}}"] = struct{}{}
			db = db.Joins(fmt.Sprintf("LEFT JOIN %[1]s{{$m2mTableName}}%[1]s ON %[1]s{{$m2mTableName}}%[1]s.%[1]s{{$object.Name | snakecase}}_{{$root.PrimaryKeyOfObject $object.Name}}%[1]s = %[2]s.%[1]s{{$root.PrimaryKeyOfObject $object.Name}}%[1]s JOIN %[1]s%[3]s%[1]s ON %[1]s{{$m2mTableName}}%[1]s.%[1]s{{$entity.GqlTypeName | snakecase}}_{{$root.PrimaryKeyOfObject $entity.GqlTypeName | snakecase}}%[1]s = %[1]s%[3]s%[1]s.%[1]s{{$root.PrimaryKeyOfObject $object.Name}}%[1]s",runtimehelper.GetQuoteChar(db), alias, tableName))
    }
		res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, fmt.Sprintf("%[1]s%[2]s%[1]s",runtimehelper.GetQuoteChar(db), tableName),true,blackList)...)
			{{- else if eq $object.Name $entity.GqlTypeName}}
		res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, fmt.Sprintf("%[1]s{{$entityGoName}}%[1]s",runtimehelper.GetQuoteChar(db)),true,blackList)...)
			{{- else }}
		if _, ok := blackList["{{$entityGoName}}"]; !ok {
			blackList["{{$entityGoName}}"] = struct{}{}
			if deep {
				tableName := db.Config.NamingStrategy.TableName("{{$root.GetGoFieldTypeName $objectName $entity }}")
				{{- $fnKeyResult := $root.ForeignName $object $entity}}
				foreignKeyName := "{{$fnKeyResult.Key | snakecase}}"
				{{- $fnKeyNumber := 5}}
				{{- $currentKeyNumber := 3}}
				{{- if eq $fnKeyResult.Table.Name $object.Name }}
					{{- $fnKeyNumber = 3}}
					{{- $currentKeyNumber = 5}}
				{{- end}}
				db = db.Joins(fmt.Sprintf("LEFT JOIN %[1]s%[2]s%[1]s %[1]s{{$entityGoName}}%[1]s ON %[1]s{{$entityGoName}}%[1]s.%[1]s%[{{$fnKeyNumber}}]s%[1]s = %[4]s.%[1]s%[{{$currentKeyNumber}}]s%[1]s",runtimehelper.GetQuoteChar(db),tableName,d.PrimaryKeyName(), alias, foreignKeyName))

			}else {
				db = db.Joins("{{$entityGoName}}")
			}
		}	
		res = append(res, d.{{$entityGoName}}.{{$methodeName}}(db, fmt.Sprintf("%[1]s{{$entityGoName}}%[1]s",runtimehelper.GetQuoteChar(db)),true,blackList)...)
			{{- end}}
    {{- end}}
	}
	{{- end}}
  {{- end}}

	return res
}
{{- end}}
{{- end}}

func (d *StringFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
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
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot,d.Not.ExtendsDatabaseQuery(db,fieldName, true, blackList)...))
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

func (d *IntFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

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
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot,d.Not.ExtendsDatabaseQuery(db,fieldName, true,blackList)...))
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

func (d *BooleanFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
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
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true,blackList)...))
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

func (d *TimeFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

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
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot,d.Not.ExtendsDatabaseQuery(db,fieldName, true,blackList)...))
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

func (d *IDFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}


	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}


	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true,blackList)...))
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