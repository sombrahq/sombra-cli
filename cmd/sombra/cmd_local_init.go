package main

import (
	"github.com/sombrahq/sombra-cli/internal/frameworks/logger"
	"github.com/sombrahq/sombra-cli/internal/runtime"
	"os"
)

type LocalInitArgs struct {
	Template string `arg:"positional,required" help:"Git Repository to use as template"`
}

func (args *LocalInitArgs) Run() {
	rt := runtime.NewLocalInitRuntime()
	cwd, err := os.Getwd()
	if err != nil {
		logger.Panic("What local directory")
	}

	err = rt.UseCase.DoLocalInit(cwd, args.Template)
	if err != nil {
		logger.Panic("Failed to init local project")
	}
}
