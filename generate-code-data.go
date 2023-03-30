package autogql

import (
	"fmt"
	"go/types"
	"log"
	"path"
	"sort"
	"strings"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/fasibio/autogql/structure"
	"github.com/vektah/gqlparser/v2/ast"
)

type GenerateData struct {
	Data    *codegen.Data
	Handler structure.SqlBuilderHelper
}

func (db *GenerateData) HookList(suffix, prefix string) []string {
	res := make([]string, 0)
	for _, v := range db.Handler.List {
		if v.HasSqlDirective() {
			res = append(res, fmt.Sprintf("%s%s%s", suffix, v.Name(), prefix))
		}
	}
	return res
}

func (db *GenerateData) HookListMany2Many(suffix string) []string {
	res := make([]string, 0)
	for _, v := range db.Handler.List {
		m2mEntities := v.Many2ManyRefEntities()
		for _, m2me := range m2mEntities {
			res = append(res, fmt.Sprintf("%s%sRef2%ssInput", suffix, m2me.GqlTypeName(), v.Name()))
		}
	}
	return res
}

func (db *GenerateData) ModelsMigrations() string {
	res := ""
	objects := make([]*structure.Object, 0)
	for _, v := range db.Handler.List {
		if v.HasSqlDirective() {
			objects = append(objects, v)

		}
	}
	sort.Slice(objects[:], func(i, j int) bool {
		return objects[i].GetOrder() < objects[j].GetOrder()
	})
	for _, o := range objects {
		splits := strings.Split(db.Data.Config.Models[o.Name()].Model[0], "/")
		res += fmt.Sprintf("&%s{},", splits[len(splits)-1])
	}

	return res[:len(res)-1]
}

func (db *GenerateData) GetPackage(v structure.Object) string {
	splits := strings.Split(db.Data.Config.Models[v.Name()].Model[0], "/")
	return strings.Split(splits[len(splits)-1], ".")[0]
}

func (db *GenerateData) PointerStrIfNeeded(typeName string, v structure.Entity, revert bool) string {
	objName := db.GetGoFieldType(typeName, v, false)
	pointerSymbol := "*"
	noSymbol := ""
	if objName[0] == '*' {
		if revert {
			return noSymbol
		}
		return pointerSymbol
	}
	if revert {
		return pointerSymbol
	}
	return noSymbol
}

func (db *GenerateData) GenPointerStrIfNeeded(typeName string, v structure.Entity, revert bool) string {
	objName := db.GetGoFieldType(typeName, v, false)
	pointerSymbol := "&"
	noSymbol := ""
	if objName[0] == '*' || objName[0:3] == "[]*" {
		if revert {
			return noSymbol
		}
		return pointerSymbol
	}
	if revert {
		return pointerSymbol
	}
	return noSymbol
}

func (db *GenerateData) GetGoFieldType(typeName string, v structure.Entity, rootType bool) string {
	objects := make(codegen.Objects, len(db.Data.Objects)+len(db.Data.Inputs))
	copy(objects, db.Data.Objects)
	copy(objects[len(db.Data.Objects):], db.Data.Inputs)
	for _, v1 := range objects {
		if v1.Name == typeName {
			fields, err := db.generateFields(v1.Definition)
			if err != nil {
				log.Panic(err)
			}

			for _, fv := range fields {
				if fv.Name == v.Name() {
					if rootType {
						return strings.TrimLeft(fv.Type.String(), "*")
					}
					return fv.Type.String()
				}
			}

		}
	}
	return templates.UcFirst(v.Name())

}

func (db *GenerateData) GetMaxMatchGoFieldType(objectname string, entities []structure.Entity) string {
	possibleRes := make(map[string]string)
	for _, e := range entities {
		tmp := db.GetGoFieldType(objectname, e, false)
		possibleRes[tmp] = tmp
	}
	if len(possibleRes) > 1 {
		return "any"
	}
	for _, k := range possibleRes {
		return k // hack to get first element
	}
	return "" // should never happend
}

func (m *GenerateData) generateFields(schemaType *ast.Definition) ([]*modelgen.Field, error) {
	cfg := m.Data.Config
	binder := cfg.NewBinder()
	fields := make([]*modelgen.Field, 0)

	for _, field := range schemaType.Fields {
		var typ types.Type
		fieldDef := cfg.Schema.Types[field.Type.Name()]

		if cfg.Models.UserDefined(field.Type.Name()) {
			var err error
			typ, err = binder.FindTypeFromName(cfg.Models[field.Type.Name()].Model[0])
			if err != nil {
				return nil, err
			}
		} else {
			switch fieldDef.Kind {
			case ast.Scalar:
				// no user defined model, referencing a default scalar
				typ = types.NewNamed(
					types.NewTypeName(0, cfg.Model.Pkg(), "string", nil),
					nil,
					nil,
				)

			case ast.Interface, ast.Union:
				// no user defined model, referencing a generated interface type
				typ = types.NewNamed(
					types.NewTypeName(0, cfg.Model.Pkg(), templates.ToGo(field.Type.Name()), nil),
					types.NewInterfaceType([]*types.Func{}, []types.Type{}),
					nil,
				)

			case ast.Enum:
				// no user defined model, must reference a generated enum
				typ = types.NewNamed(
					types.NewTypeName(0, cfg.Model.Pkg(), templates.ToGo(field.Type.Name()), nil),
					nil,
					nil,
				)

			case ast.Object, ast.InputObject:
				// no user defined model, must reference a generated struct
				typ = types.NewNamed(
					types.NewTypeName(0, cfg.Model.Pkg(), templates.ToGo(field.Type.Name()), nil),
					types.NewStruct(nil, nil),
					nil,
				)

			default:
				panic(fmt.Errorf("unknown ast type %s", fieldDef.Kind))
			}
		}

		name := templates.ToGo(field.Name)
		if nameOveride := cfg.Models[schemaType.Name].Fields[field.Name].FieldName; nameOveride != "" {
			name = nameOveride
		}

		typ = binder.CopyModifiersFromAst(field.Type, typ)

		if cfg.StructFieldsAlwaysPointers {
			if isStruct(typ) && (fieldDef.Kind == ast.Object || fieldDef.Kind == ast.InputObject) {
				typ = types.NewPointer(typ)
			}
		}

		f := &modelgen.Field{
			Name:        field.Name,
			GoName:      name,
			Type:        typ,
			Description: field.Description,
			Tag:         `json:"` + field.Name + `"`,
		}
		fields = append(fields, f)
	}

	return fields, nil
}

func (db *GenerateData) GetGoFieldTypeName(typeName string, v structure.Entity) string {
	objects := make(codegen.Objects, len(db.Data.Objects)+len(db.Data.Inputs))
	copy(objects, db.Data.Objects)
	copy(objects[len(db.Data.Objects):], db.Data.Inputs)
	for _, v1 := range objects {
		if v1.Name == typeName {
			for _, fv := range v1.Fields {
				if fv.Name == v.Name() {
					return fv.Type.Name()
				}
			}

		}
	}
	return templates.UcFirst(v.Name())
}

func (db *GenerateData) GetGoFieldName(typeName string, v structure.Entity) string {
	objects := make(codegen.Objects, len(db.Data.Objects)+len(db.Data.Inputs))
	copy(objects, db.Data.Objects)
	copy(objects[len(db.Data.Objects):], db.Data.Inputs)
	for _, v1 := range objects {
		if v1.Name == typeName {
			for _, fv := range v1.Fields {
				if fv.Name == v.Name() {
					return fv.GoFieldName
				}
			}

		}
	}
	return templates.UcFirst(v.Name())
}

func (db *GenerateData) GetPointerSymbol(entity structure.Entity) string {
	if entity.IsPrimitive() {
		return ""
	}
	if entity.IsArray() {
		return ""
	}
	return "&"
}

func (db *GenerateData) GetValueOfInput(objectname string, builder structure.Object, v structure.Entity) string {
	if v.IsPrimitive() {
		refSymbol := ""
		if !v.Required() {
			refSymbol = "*"
		}
		return fmt.Sprintf("%s%s.%s", refSymbol, objectname, db.GetGoFieldName(objectname, v))
	}
	// is a ref
	refSign := "&"
	if v.IsArray() {
		refSign = ""
	}
	return fmt.Sprintf("%s%s", refSign, v.Name())
}

func (db *GenerateData) ForeignName(object structure.Object, entity structure.Entity) string {
	d, ok := db.Handler.List[entity.GqlTypeName()]
	if !ok {
		panic("ForeignName: Can not find object " + entity.GqlTypeName())
	}
	res := strings.ToLower(d.PrimaryKeyField().Name())
	if res == "" {
		res = object.ForeignNameKeyName(strings.ToLower(entity.Name()))
	}
	return res
}

func (db *GenerateData) PrimaryKeyOfObject(o string) string {
	return db.PrimaryKeyEntityOfObject(o).Name()
}

func (db *GenerateData) PrimaryKeyEntityOfObject(o string) *structure.Entity {
	d, ok := db.Handler.List[o]
	if !ok {
		panic("PrimaryKeyEntityOfObject: Can not find object " + o)
	}
	return d.PrimaryKeyField()
}

func (db *GenerateData) Imports() []string {
	var addedImports []string
	seenImports := make(map[string]bool)

	for _, v := range db.Handler.List {
		sp := strings.LastIndex(db.Data.Config.Models[v.Name()].Model[0], ".")
		importName := db.Data.Config.Models[v.Name()].Model[0][:sp]
		seenImports[importName] = true
	}
	addedImports = append(addedImports, db.Data.Config.Exec.ImportPath())
	for key := range seenImports {
		addedImports = append(addedImports, key)
	}
	addedImports = append(addedImports, "github.com/fasibio/autogql/runtimehelper")
	addedImports = append(addedImports, path.Join(db.Data.Config.Resolver.ImportPath(), "db"))
	return addedImports
}

func (db *GenerateData) GeneratedPackage() string {

	if db.Data.Config.Resolver.Package == db.Data.Config.Exec.Package {
		return ""
	}
	return db.Data.Config.Exec.Package + "."
}

func isStruct(t types.Type) bool {
	_, is := t.Underlying().(*types.Struct)
	return is
}
