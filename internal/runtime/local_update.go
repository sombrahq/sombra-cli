package runtime

import (
	"github.com/sombrahq/sombra-cli/internal/core/usecases"
	"github.com/sombrahq/sombra-cli/internal/frameworks/cvs"
	"github.com/sombrahq/sombra-cli/internal/frameworks/files"
	"github.com/sombrahq/sombra-cli/internal/frameworks/sombra"
	"github.com/sombrahq/sombra-cli/internal/frameworks/templates"
	"github.com/sombrahq/sombra-cli/internal/frameworks/versions"
)

type LocalUpdateRuntime struct {
	UseCase usecases.CliUpdateCase
}

func NewLocalUpdateRuntime() *LocalUpdateRuntime {
	dirManager := files.NewDirectoryScannerService()
	fileManager := files.NewFileManagerService()
	stringProcessor := sombra.NewProcessor()
	engine := usecases.NewSombraEngineInteractor(dirManager, fileManager, stringProcessor)
	templateDef := templates.NewDefService()

	repoPrepare := usecases.NewRepositoryPrepareInteractor(cvs.For)
	sombraDefManager := sombra.NewDefService()
	versionManager := versions.NewTemplateTagManagerService()

	patchManager := cvs.NewPatchService()

	copyCase := usecases.NewLocalCopyInteractor(repoPrepare, templateDef, sombraDefManager, versionManager, dirManager, fileManager, engine)
	diffCase := usecases.NewDirectoryLocalDiffInteractor(repoPrepare, patchManager, templateDef, sombraDefManager, versionManager, dirManager, fileManager, engine)
	cliCase := usecases.NewCliUpdateInteractor(copyCase, diffCase)
	return &LocalUpdateRuntime{
		UseCase: cliCase,
	}
}
