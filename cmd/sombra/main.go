package main

import (
	"github.com/alexflint/go-arg"
	"github.com/sombrahq/sombra-cli/internal/frameworks/logger"
)

/***********
COMMANDS
************/

var args struct {
	Local    *LocalSubcommand    `arg:"subcommand:local"`
	Template *TemplateSubcommand `arg:"subcommand:template"`
}

/***********
CONFIG
************/

func main() {
	logger.Init()
	arg.MustParse(&args)

	switch {
	case args.Local != nil:
		args.Local.Run()
	case args.Template != nil:
		args.Template.Run()
	default:
		logger.Panic("No command specified")
	}
}
