package runtime

import (
	"github.com/sombrahq/sombra-cli/internal/core/usecases"
	"github.com/sombrahq/sombra-cli/internal/frameworks/analysers"
	"github.com/sombrahq/sombra-cli/internal/frameworks/files"
	"github.com/sombrahq/sombra-cli/internal/frameworks/sombra"
	"github.com/sombrahq/sombra-cli/internal/frameworks/templates"
)

type TemplateInitRuntime struct {
	UseCase usecases.CliTemplateInitCase
}

func NewTemplateRuntime() *TemplateInitRuntime {
	dirManager := files.NewDirectoryScannerService()
	fileManager := files.NewFileManagerService()
	stringProcessor := sombra.NewProcessor()
	engine := usecases.NewSombraEngineInteractor(dirManager, fileManager, stringProcessor)
	registry := analysers.GetRegistry()
	templateDef := templates.NewDefService()

	directoryInit := usecases.NewDirectoryTemplateInitInteractor(dirManager, registry, templateDef, engine)
	cliCase := usecases.NewCliTemplateInitInteractor(directoryInit)
	return &TemplateInitRuntime{
		UseCase: cliCase,
	}
}
