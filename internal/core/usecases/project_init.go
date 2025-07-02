package usecases

import (
	"github.com/sombrahq/sombra-cli/internal/core/entities"
)

type VariableReaderPort interface {
	GetValues(vars []string) *entities.Mappings
}

type LocalInitCase interface {
	LocalInit(target, uri string) error
}

type LocalInitInteractor struct {
	repoPrepare        RepositoryPrepareCase
	templateDefManager TemplateDefManagerPort
	sombraDefManager   SombraDefManagerPort
	varsSource         VariableReaderPort
}

func NewLocalInitInteractor(
	repoPrepare RepositoryPrepareCase,
	templateDefManager TemplateDefManagerPort,
	sombraDefManager SombraDefManagerPort,
	varsSource VariableReaderPort,
) *LocalInitInteractor {
	return &LocalInitInteractor{
		repoPrepare:        repoPrepare,
		templateDefManager: templateDefManager,
		sombraDefManager:   sombraDefManager,
		varsSource:         varsSource,
	}
}

func (l *LocalInitInteractor) LocalInit(target, uri string) error {
	// Download and prepare the version
	repo, err := l.repoPrepare.Prepare(uri, "")
	if err != nil {
		return err
	}
	defer repo.Clean()

	// Read the template configuration
	fn := l.templateDefManager.GetFile(repo.Dir())
	tpl, err := l.templateDefManager.Load(fn)
	if err != nil {
		return err
	}

	// Read missing variables
	vars := tpl.Vars
	mappings := l.varsSource.GetValues(vars)

	// Read sombra file
	fn = l.sombraDefManager.GetFile(target)
	def, err := l.sombraDefManager.Load(fn)
	if err != nil {
		return err
	}

	// Update sombra file
	def.Templates = append(def.Templates, &entities.TemplateConfig{
		URI:  uri,
		Vars: *mappings,
	})

	// Store sombra file
	err = l.sombraDefManager.Save(fn, def)
	if err != nil {
		return err
	}

	return nil
}

var _ LocalInitCase = (*LocalInitInteractor)(nil)
