package autogql

import (
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/fasibio/autogql/structure"
)

type AutoGqlPlugin struct {
	Handler structure.SqlBuilderHelper
}

func NewAutoGqlPlugin() (*AutoGqlPlugin, *modelgen.Plugin) {
	sp := &AutoGqlPlugin{}
	modelGenPlugin := &modelgen.Plugin{MutateHook: MutateHook(sp), FieldHook: ConstraintFieldHook(sp)}
	return sp, modelGenPlugin
}

func (ggs *AutoGqlPlugin) Name() string {
	return "autogql"
}
