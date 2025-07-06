package usecases

import (
	"github.com/sombrahq/sombra-cli/internal/core/entities"
	"path/filepath"
)

type LocalCopyInteractor struct {
	repoPrepare        RepositoryPrepareCase
	templateDefManager TemplateDefManagerPort
	sombraDefManager   SombraDefManagerPort
	versionManager     VersionManagerPort
	scanner            DirectoryManagerPort
	localFiles         FileManagerPort
	engine             SombraEngineCase
}

func NewLocalCopyInteractor(
	repoPrepare RepositoryPrepareCase,
	templateDefManager TemplateDefManagerPort,
	sombraDefManager SombraDefManagerPort,
	versionManager VersionManagerPort,
	scanner DirectoryManagerPort,
	localFiles FileManagerPort,
	engine SombraEngineCase,
) *LocalCopyInteractor {
	return &LocalCopyInteractor{
		repoPrepare:        repoPrepare,
		templateDefManager: templateDefManager,
		sombraDefManager:   sombraDefManager,
		versionManager:     versionManager,
		scanner:            scanner,
		localFiles:         localFiles,
		engine:             engine,
	}
}

func (copy *LocalCopyInteractor) LocalUpdate(target, uri, tag string) error {
	// Read sombra file
	sombraFile := copy.sombraDefManager.GetFile(target)
	def, err := copy.sombraDefManager.Load(sombraFile)
	if err != nil {
		return err
	}

	// Download and prepare the version
	repo, err := copy.repoPrepare.Prepare(uri, "")
	if err != nil {
		return err
	}
	defer repo.Clean()

	// If the tag variable is empty, find the latest tag
	var version entities.Version
	var tags []string
	if tag == "" {
		tags, err = repo.GetTags()
		if err != nil {
			return err
		}
		version, err = copy.versionManager.GetLatest(tags, "*")
		if err != nil {
			return err
		}
	} else {
		version = entities.Version(tag)
	}

	// Iterate over all templates
	var tpl *entities.TemplateDef
	var fn entities.File
	for _, template := range def.Templates {
		if template.URI != uri {
			continue
		}

		// Render TemplateConfig Definition using Sombra configuration
		fn = copy.templateDefManager.GetFile(repo.Dir())
		tpl, err = copy.templateDefManager.Render(fn, template.Vars)
		if err != nil {
			return err
		}

		// Execute the mappings
		err = copy.copyFiles(repo.Dir(), filepath.Join(target, template.Path), tpl)
		if err != nil {
			return err
		}

		// Update the template configuration
		template.Current = version
	}

	// Store sombra file
	err = copy.sombraDefManager.Save(sombraFile, def)
	if err != nil {
		return err
	}

	return nil
}

func (copy *LocalCopyInteractor) copyFiles(templateDir, targetDir string, templateConfig *entities.TemplateDef) error {
	tree := copy.scanner.ScanTree(templateDir, []entities.Wildcard{"**/*"}, nil)
	var fn entities.File
	var items *entities.MapResult
	for result := range tree {
		if result.Err != nil {
			return result.Err
		}

		fn = result.File

		match, res, err := copy.engine.Match(fn, templateConfig.Patterns)
		if err != nil {
			return err
		}

		// Notice that match can be false even if res is not empty
		// This is because the mapping engine only matches a file
		// if there are non-abstract mappings matching the file
		if !match {
			continue
		}

		items = copy.engine.Combine(res)

		if result.IsDir {
			err = copy.processDir(targetDir, fn, items)
		} else {
			err = copy.processFile(templateDir, targetDir, fn, items)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (copy *LocalCopyInteractor) processDir(target string, path entities.File, res *entities.MapResult) error {
	newDir := copy.engine.NewFile(path, res.Path, res.Name)
	return copy.localFiles.EnsureDir(target, newDir)
}

func (copy *LocalCopyInteractor) processFile(src string, target string, file entities.File, res *entities.MapResult) error {
	newFile := copy.engine.NewFile(file, res.Path, res.Name)
	content, err := copy.localFiles.Read(src, file)
	if err != nil {
		return err
	}

	newContent := copy.engine.NewContent(content, res.Content)
	err = copy.localFiles.Write(target, newFile, newContent)
	return err
}

var _ LocalUpdateCase = (*LocalCopyInteractor)(nil)
