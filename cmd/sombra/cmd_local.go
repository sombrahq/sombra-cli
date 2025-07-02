package main

import (
	"github.com/sombrahq/sombra-cli/internal/frameworks/logger"
)

type LocalSubcommand struct {
	LocalInit   *LocalInitArgs   `arg:"subcommand:init"`
	LocalUpdate *LocalUpdateArgs `arg:"subcommand:update"`
}

func (args *LocalSubcommand) Run() {
	switch {
	case args.LocalInit != nil:
		args.LocalInit.Run()
	case args.LocalUpdate != nil:
		args.LocalUpdate.Run()

	default:
		logger.Panic("command not supported")
	}

}
