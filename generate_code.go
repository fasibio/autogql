package autogql

import (
	_ "embed"
	"log"
	"path"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/Masterminds/sprig/v3"
)

func (ggs *AutoGqlPlugin) GenerateCode(data *codegen.Data) error {
	err := ggs.generateResolverCode(data)
	if err != nil {
		return err
	}
	err = ggs.generateExtraFunctionsToModels(data)
	if err != nil {
		return err
	}
	err = ggs.generateFilterModelFunctions(data)
	if err != nil {
		return err
	}
	return ggs.generateDbCode(data)
}

//go:embed generate_code_resolver.go.tpl
var generateResolverCodeTemplate string

//go:embed generate_code_db.go.tpl
var generateDbCodeTemplate string

//go:embed generate_code_model.go.tpl
var generateModelCodeTemplate string

//go:embed generate_code_filter_model.go.tpl
var generateFilterModelCodeTemplate string

func (ggs AutoGqlPlugin) generateDbCode(data *codegen.Data) error {
	filename := path.Join(data.Config.Resolver.Package, "db/db_gen.go")
	log.Println("generateDbCode", filename)

	return templates.Render(templates.Options{
		PackageName: "db",
		Filename:    filename,
		Data: &GenerateData{
			Data:    data,
			Handler: ggs.Handler,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        generateDbCodeTemplate,
		Funcs:           sprig.FuncMap(),
	})
}

func (ggs AutoGqlPlugin) generateExtraFunctionsToModels(data *codegen.Data) error {
	filename := path.Join(data.Config.Resolver.Package, data.Config.Model.Package, "models_autogql.go")
	log.Println("generateExtraFunctionsToModels", filename)

	return templates.Render(templates.Options{
		PackageName: "model",
		Filename:    filename,
		Data: &GenerateData{
			Data:    data,
			Handler: ggs.Handler,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        generateModelCodeTemplate,
		Funcs:           sprig.FuncMap(),
	})

}

func (ggs AutoGqlPlugin) generateFilterModelFunctions(data *codegen.Data) error {
	filename := path.Join(data.Config.Resolver.Package, data.Config.Model.Package, "models_filter_autogql.go")
	log.Println("generateFilterModelFunctions", filename)

	return templates.Render(templates.Options{
		PackageName: "model",
		Filename:    filename,
		Data: &GenerateData{
			Data:    data,
			Handler: ggs.Handler,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        generateFilterModelCodeTemplate,
		Funcs:           sprig.FuncMap(),
	})

}

func (ggs *AutoGqlPlugin) generateResolverCode(data *codegen.Data) error {
	filename := path.Join(data.Config.Resolver.Package, "autogql.resolvers.go")
	log.Println("generateResolverCode", filename)

	return templates.Render(templates.Options{
		PackageName: "graph",
		Filename:    filename,
		Data: &GenerateData{
			Data:    data,
			Handler: ggs.Handler,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        generateResolverCodeTemplate,
		Funcs:           sprig.FuncMap(),
	})
}
