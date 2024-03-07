package autogql

import (
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/fasibio/autogql/structure"
)

type AutoGqlPlugin struct {
	Handler structure.SqlBuilderHelper
}

func NewAutoGqlPlugin(cfg *config.Config) (plugin.Plugin, *modelgen.Plugin) {
	cfg.Models.Add("SoftDelete", "github.com/fasibio/autogql/runtimehelper.SoftDelete")
	sp := &AutoGqlPlugin{}
	modelGenPlugin := &modelgen.Plugin{MutateHook: MutateHook(sp), FieldHook: ConstraintFieldHook(sp)}
	return sp, modelGenPlugin
}

func (ggs *AutoGqlPlugin) Name() string {
	return "autogql"
}
