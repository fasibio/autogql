package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/fasibio/autogql"
)

func main() {
	cfg, err := config.LoadConfig("gqlgen.yml")

	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	sqlPlugin, muateHookPlugin := autogql.NewAutoGqlPlugin()
	err = api.Generate(cfg, api.AddPlugin(sqlPlugin), api.ReplacePlugin(muateHookPlugin))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
