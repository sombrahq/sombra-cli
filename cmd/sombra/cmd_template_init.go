package main

import (
	"github.com/sombrahq/sombra-cli/internal/frameworks/logger"
	"github.com/sombrahq/sombra-cli/internal/runtime"
)

type TemplateInitArgs struct {
	Dir     string   `arg:"positional" default:"." help:"Directory of the project to initialize as template"`
	Exclude []string `arg:"-e,--exclude,separate" help:"Wildcard of the files to exclude"`
	Only    []string `arg:"-o,--only,separate"  help:"Wildcard of the files to include"`
}

func (args *TemplateInitArgs) Run() {
	rt := runtime.NewTemplateRuntime()

	if args.Exclude == nil {
		args.Exclude = []string{}
	}

	if args.Only == nil {
		args.Only = []string{"/**/*"}
	}

	err := rt.UseCase.DoTemplateInit(args.Dir, args.Only, args.Exclude)
	if err != nil {
		logger.Panic("Failed to init template")
	}
}
