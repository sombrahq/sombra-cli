package main

import (
	"github.com/sombrahq/sombra-cli/internal/frameworks/logger"
)

type TemplateSubcommand struct {
	TemplateInit *TemplateInitArgs `arg:"subcommand:init"`
}

func (args *TemplateSubcommand) Run() {
	switch {
	case args.TemplateInit != nil:
		args.TemplateInit.Run()

	default:
		logger.Panic("command not supported")
	}

}
