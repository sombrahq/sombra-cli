package usecases

import "github.com/sombrahq/sombra-cli/internal/core/entities"

type TemplateDefManagerPort interface {
	GetFile(dir string) entities.File
	Load(def entities.File) (*entities.TemplateDef, error)
	Save(def entities.File, templateDef *entities.TemplateDef) error
	Render(def entities.File, vars entities.Mappings) (*entities.TemplateDef, error)
}
