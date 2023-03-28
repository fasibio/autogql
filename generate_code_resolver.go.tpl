
{{ reserveImport "context"  }}
{{ reserveImport "fmt"  }}
{{ reserveImport "gorm.io/gorm/clause"}}
{{ range $import := .Imports }}
	{{ reserveImport $import }}
{{end}}

{{$root := .}}
{{$hookBaseName := "AutoGql"}}

{{- range $objectName, $object := .Handler.List.Objects}}
{{- if $object.HasSqlDirective}}
  {{- if $object.SQLDirective.HasQueries}}
    {{- if $object.SQLDirective.Query.Get}}
// Get{{$object.Name}} is the resolver for the get{{$object.Name}} field.
{{- $primaryFields := $object.PrimaryKeys }}
func (r *queryResolver) Get{{$object.Name}}(ctx context.Context, {{range $primaryFieldKey, $primaryField := $primaryFields}} {{$primaryField.Name}} {{$root.GetGoFieldType $objectName $primaryField false}}, {{end }}) (*model.{{$object.Name}}, error) {
	v, okHook := r.Sql.Hooks[string(db.Get{{$object.Name}})].(db.{{$hookBaseName}}HookGet[model.{{$object.Name}}, {{$root.GetMaxMatchGoFieldType $objectName $primaryFields}}])
	db := r.Sql.Db
	if okHook {
		var err error
		db, err = v.Received(ctx, r.Sql, {{range $primaryFieldKey, $primaryField := $primaryFields}} {{$primaryField.Name}}, {{end }})
		if err != nil {
			return nil, err
		}
	}
	db = runtimehelper.GetPreloadSelection(ctx,db,runtimehelper.GetPreloadsMap(ctx,"{{$object.Name}}"))
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	var res model.{{$object.Name}}
  db = db.First(&res, {{range $primaryFieldKey, $primaryField := $primaryFields}} {{$primaryField.Name}}, {{end }})
	if okHook {
		r, err := v.AfterCallDb(ctx, &res)
		if err != nil {
			return nil, err
		}
		res = *r
		r, err = v.BeforeReturn(ctx, &res, db)
		if err != nil {
			return nil, err
		}
		res = *r
	}
	return &res, db.Error
}
    {{- end}}
    {{- if $object.SQLDirective.Query.Query}}
// Query{{$object.Name}} is the resolver for the query{{$object.Name}} field.
func (r *queryResolver) Query{{$object.Name}}(ctx context.Context, filter *model.{{$object.Name}}FiltersInput, order *model.{{$object.Name}}Order, first *int, offset *int) (*model.{{$object.Name}}QueryResult, error) {
	v, okHook := r.Sql.Hooks[string(db.Query{{$object.Name}})].(db.{{$hookBaseName}}HookQuery[model.{{$object.Name}}, model.{{$object.Name}}FiltersInput,model.{{$object.Name}}Order])
  db := r.Sql.Db
	if okHook {
		var err error
		db,filter,order,first, offset, err = v.Received(ctx,r.Sql,filter,order,first,offset)
		if err != nil {
			return nil, err
		}
	}
	var res []*model.{{$object.Name}}
  tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	db = runtimehelper.GetPreloadSelection(ctx, db,runtimehelper.GetPreloadsMap(ctx, "data").SubTables[0])
	if filter != nil{
		blackList := make(map[string]struct{})
		sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(db, tableName, false, blackList), "AND")
		db.Where(sql, arguments...)
	}

	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx,db)
		if err != nil {
			return nil, err
		}
	}

	if (order != nil){
		if order.Asc != nil {
			db = db.Order(fmt.Sprintf("%s.%s asc",tableName,order.Asc))
		}
		if order.Desc != nil {
			db = db.Order(fmt.Sprintf("%s.%s desc",tableName,order.Desc))
		}
	}
	var total int64
	db.Model(res).Count(&total)
	if first != nil {
		db = db.Limit(*first)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	db = db.Find(&res)
	if okHook {
		var err error
		res, err = v.AfterCallDb(ctx,res)
		if err != nil {
			return nil, err
		}
		res, err = v.BeforeReturn(ctx,res,db)
		if err != nil {
			return nil, err
		}
	}
	return &model.{{$object.Name}}QueryResult{
		Data: res,
    Count: len(res),
		TotalCount: int(total),
	},db.Error
}
    {{- end}}
  {{- end}}
  {{- if $object.SQLDirective.HasMutation}}
func (r *Resolver) Add{{$object.Name}}Payload() {{$root.GeneratedPackage}}Add{{$object.Name}}PayloadResolver { return &{{lcFirst $object.Name}}PayloadResolver[*model.Add{{$object.Name}}Payload]{r} }
func (r *Resolver) Delete{{$object.Name}}Payload() {{$root.GeneratedPackage}}Delete{{$object.Name}}PayloadResolver { return &{{lcFirst $object.Name}}PayloadResolver[*model.Delete{{$object.Name}}Payload]{r} }
func (r *Resolver) Update{{$object.Name}}Payload() {{$root.GeneratedPackage}}Update{{$object.Name}}PayloadResolver { return &{{lcFirst $object.Name}}PayloadResolver[*model.Update{{$object.Name}}Payload]{r} }


type {{lcFirst $object.Name}}Payload interface {
	*model.Add{{$object.Name}}Payload | *model.Delete{{$object.Name}}Payload | *model.Update{{$object.Name}}Payload
}

type {{lcFirst $object.Name}}PayloadResolver[T  {{lcFirst $object.Name}}Payload] struct {
	*Resolver
}
func (r *{{lcFirst $object.Name}}PayloadResolver[T]) {{$object.Name}}(ctx context.Context, obj T, filter *model.{{$object.Name}}FiltersInput, order *model.{{$object.Name}}Order, first *int, offset *int) (*model.{{$object.Name}}QueryResult, error){
	return r.Query().Query{{$object.Name}}(ctx,filter,order,first,offset)
}
		{{- range $m2mKey, $m2mEntity := $object.Many2ManyRefEntities }}
func (r *mutationResolver) Add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s(ctx context.Context, input model.{{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput) (*model.Update{{$object.Name}}Payload, error){
	v, okHook := r.Sql.Hooks[string(db.Add{{$m2mEntity.GqlTypeName}}2{{$object.Name}}s)].(db.{{$hookBaseName}}HookMany2Many[model.{{$m2mEntity.GqlTypeName}}Ref2{{$object.Name}}sInput,model.Update{{$object.Name}}Payload])
	db := r.Sql.Db
	if okHook {
		var err error
		db, input, err = v.Received(ctx,r.Sql,&input)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	blackList := make(map[string]struct{})
	sql, arguments := runtimehelper.CombineSimpleQuery(input.Filter.ExtendsDatabaseQuery(r.Sql.Db, tableName, false, blackList), "AND")
	db = db.Model(&model.{{$object.Name}}{}).Where(sql, arguments...)
	var res []*model.{{$object.Name}}
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db )	
		if err != nil {
			return nil, err
		}
	}
	db.Find(&res)
	{{- $table1ID := $root.GetGoFieldName $object.Name $object.PrimaryKeyField }}
	{{- $tabe2PrimaryEntity := $root.PrimaryKeyEntityOfObject $m2mEntity.GqlTypeName}}
	{{- $table2ID := $root.GetGoFieldName $m2mEntity.GqlTypeName $tabe2PrimaryEntity }}
	type {{camelcase $m2mKey}} struct {
		{{ucFirst $object.Name}}{{$table1ID}} {{$root.GetGoFieldType $object.Name $object.PrimaryKeyField false}} 
		{{ucFirst $m2mEntity.GqlTypeName}}{{$table2ID}} {{$root.GetGoFieldType $m2mEntity.GqlTypeName $tabe2PrimaryEntity false}} 
	}
	resIds := make([]map[string]interface{}, 0)
	for _, v := range res{
		for _, v1 := range input.Set {
				tmp := make(map[string]interface{})
				tmp["{{ucFirst $object.Name}}{{$table1ID}}"] = v.ID
				tmp["{{ucFirst $m2mEntity.GqlTypeName}}{{$table2ID}}"] = v1
				resIds = append(resIds, tmp)
		}	
	}
	d := r.Sql.Db.Model(&{{camelcase $m2mKey}}{}).Create(resIds)
	result :=  &model.Update{{$object.Name}}Payload{
		Count: int(d.RowsAffected),
	}
	if okHook {
		var err error
		result, err =v.BeforeReturn(ctx,db, result)
		if err != nil {
			return nil, err
		}
	}
	return result,d.Error
}
		{{- end}}
    {{- if $object.SQLDirective.Mutation.Add}}
// Add{{$object.Name}} is the resolver for the add{{$object.Name}} field.
func (r *mutationResolver) Add{{$object.Name}}(ctx context.Context, input []*model.{{$object.Name}}Input) (*model.Add{{$object.Name}}Payload, error) {
	v, okHook := r.Sql.Hooks[string(db.Add{{$object.Name}})].(db.{{$hookBaseName}}HookAdd[model.{{$object.Name}}, model.{{$object.Name}}Input, model.Add{{$object.Name}}Payload])
	res := &model.Add{{$object.Name}}Payload{}
	db := r.Sql.Db
	if okHook {
		var err error
		db,input, err = v.Received(ctx,r.Sql,input)
		if err != nil {
			return nil, err
		}
	}
	obj:= make([]model.{{$object.Name}}, len(input))
  for i, v := range input {
    obj[i] = v.MergeToType()
  }
	db = db.Omit(clause.Associations)
	if okHook {
		var err error
		db, obj, err = v.BeforeCallDb(ctx,db,obj)
		if err != nil {
			return nil, err
		}
	}
  db = db.Create(&obj)
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx,db, obj, res)
		if err != nil {
			return nil, err
		}
	}
  return res, db.Error
}
    {{- end}}
    {{- if $object.SQLDirective.Mutation.Update}}
// Update{{$object.Name}} is the resolver for the update{{$object.Name}} field.
func (r *mutationResolver) Update{{$object.Name}}(ctx context.Context, input model.Update{{$object.Name}}Input) (*model.Update{{$object.Name}}Payload, error) {
  v, okHook := r.Sql.Hooks[string(db.Update{{$object.Name}})].(db.{{$hookBaseName}}HookUpdate[ model.Update{{$object.Name}}Input, model.Update{{$object.Name}}Payload])
	db := r.Sql.Db
	if okHook{
		var err error
		db, input, err =v.Received(ctx,r.Sql,&input)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName+".{{$root.PrimaryKeyOfObject $object.Name}}")
	sql, arguments := runtimehelper.CombineSimpleQuery(input.Filter.ExtendsDatabaseQuery(queryDb, tableName, false, blackList), "AND")
	obj := model.{{$object.Name}}{}
  queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.{{$object.Name}}
	queryDb.Find(&toChange)
	update := input.Set.MergeToType()
	if okHook {
		var err error
		db, update, err = v.BeforeCallDb(ctx,db,update)
		if err != nil {
			return nil, err
		}
	}
	{{- $primaryEntity := $root.PrimaryKeyEntityOfObject $object.Name}}
	ids := make([]{{$root.GetGoFieldType $object.Name $primaryEntity false}},len(toChange))
	for i, one := range toChange {
		ids[i] = one.{{$root.GetGoFieldTypeName $object.Name $primaryEntity}}
	}
	db = db.Model(&obj).Where("{{$root.PrimaryKeyOfObject $object.Name}} IN ?",ids).Updates(update)
	res := &model.Update{{$object.Name}}Payload{
		Count: int(db.RowsAffected),
	}
		if okHook {
		var err error 
		res, err = v.BeforeReturn(ctx, db, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}
    {{- end}}
    {{- if $object.SQLDirective.Mutation.Delete}}
// Delete{{$object.Name}} is the resolver for the delete{{$object.Name}} field.
func (r *mutationResolver) Delete{{$object.Name}}(ctx context.Context, filter model.{{$object.Name}}FiltersInput) (*model.Delete{{$object.Name}}Payload, error) {
	v, okHook := r.Sql.Hooks[string(db.Delete{{$object.Name}})].(db.{{$hookBaseName}}HookDelete[model.{{$object.Name}}FiltersInput, model.Delete{{$object.Name}}Payload])
	db := r.Sql.Db
	if okHook{
		var err error
		db, filter, err = v.Received(ctx, r.Sql, &filter)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("{{$object.Name}}")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName+".{{$root.PrimaryKeyOfObject $object.Name}}")
	sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(queryDb, tableName, false, blackList), "AND")
	obj := model.{{$object.Name}}{}
	queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.{{$object.Name}}
	queryDb.Find(&toChange)
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx,db)
		if err != nil {
			return nil, err
		}
	}
	{{- $primaryEntity := $root.PrimaryKeyEntityOfObject $object.Name}}
	ids := make([]{{$root.GetGoFieldType $object.Name $primaryEntity false}},len(toChange))
	for i, one := range toChange {
		ids[i] = one.{{$root.GetGoFieldTypeName $object.Name $primaryEntity}}
	}
	db = db.Model(&obj).Where("{{$root.PrimaryKeyOfObject $object.Name}} IN ?",ids).Delete(&obj)
		msg := fmt.Sprintf("%d rows deleted", db.RowsAffected)
	res := &model.Delete{{$object.Name}}Payload{
		Count: int(db.RowsAffected),
		Msg:   &msg,
	}
	if okHook{
		var err error
		res, err = v.BeforeReturn(ctx,db,res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}
    {{- end}}
  {{- end}}
{{- end}}
{{- end}}