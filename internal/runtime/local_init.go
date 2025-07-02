package runtime

import (
	"github.com/sombrahq/sombra-cli/internal/core/usecases"
	"github.com/sombrahq/sombra-cli/internal/frameworks/cvs"
	"github.com/sombrahq/sombra-cli/internal/frameworks/sombra"
	"github.com/sombrahq/sombra-cli/internal/frameworks/templates"
	"github.com/sombrahq/sombra-cli/internal/frameworks/vars"
)

type LocalInitRuntime struct {
	UseCase usecases.CliLocalInitCase
}

func NewLocalInitRuntime() *LocalInitRuntime {
	var repoPrepare usecases.RepositoryPrepareCase = usecases.NewRepositoryPrepareInteractor(cvs.For)
	var templateDefManager usecases.TemplateDefManagerPort = templates.NewDefService()
	var sombraDefManager usecases.SombraDefManagerPort = sombra.NewDefService()
	var varsSource = vars.NewReader()
	localInitCase := usecases.NewLocalInitInteractor(repoPrepare, templateDefManager, sombraDefManager, varsSource)
	cliCase := usecases.NewCliLocalInitInteractor(localInitCase)
	return &LocalInitRuntime{
		UseCase: cliCase,
	}
}
