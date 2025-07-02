package main

import (
	"github.com/sombrahq/sombra-cli/internal/frameworks/logger"
	"github.com/sombrahq/sombra-cli/internal/runtime"
	"os"
)

type LocalUpdateArgs struct {
	Template string `arg:"positional,required" help:"Git template to update"`
	Tag      string `arg:"--tag" help:"Git tag to use as template"`
	Method   string `arg:"--method" help:"Method to use for updating the project. (copy|diff)" default:"copy"`
}

func (args *LocalUpdateArgs) Run() {
	rt := runtime.NewLocalUpdateRuntime()
	cwd, err := os.Getwd()
	if err != nil {
		logger.Panic("What local directory")
	}

	err = rt.UseCase.DoLocalUpdate(cwd, args.Template, args.Tag, args.Method)
	if err != nil {
		logger.Panic("Failed to do dry copy")
	}
}
