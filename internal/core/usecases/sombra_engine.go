package usecases

import (
	"github.com/sombrahq/sombra-cli/internal/core/entities"
	"path/filepath"
	"sort"
)

type SombraStringsPort interface {
	ProcessString(target string, mapping entities.MapList) string
	ProcessContent(content []byte, mapping entities.MapList) []byte
}

type SombraEngineCase interface {
	Match(file entities.File, mappings []*entities.Pattern) (bool, []*entities.Pattern, error)
	Combine(patterns []*entities.Pattern) *entities.MapResult
	NewFile(file entities.File, paths entities.MapList, names entities.MapList) entities.File
	NewContent(content []byte, mappings entities.MapList) []byte
}

type SombraEngineInteractor struct {
	alwaysIgnore    []entities.Wildcard
	dirManager      DirectoryManagerPort
	fileManager     FileManagerPort
	stringProcessor SombraStringsPort
}

func (l *SombraEngineInteractor) Match(file entities.File, mappings []*entities.Pattern) (bool, []*entities.Pattern, error) {
	var match bool
	include := false
	var err error

	res := make([]*entities.Pattern, 0)

	for _, value := range mappings {
		match, err = l.dirManager.PathMatch(file, value.Pattern)
		if err != nil {
		}

		if !match {
			continue
		}

		skip := false

		exceptions := append(value.Except, l.alwaysIgnore...)
		for _, except := range exceptions {
			match, err = l.dirManager.PathMatch(file, except)
			if err != nil {
				return false, nil, err
			}
			if match {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		if !value.Abstract {
			include = true
		}
		res = append(res, value)
	}

	return include, res, nil
}

func (l *SombraEngineInteractor) Combine(patterns []*entities.Pattern) *entities.MapResult {
	path := entities.Mappings{}
	name := entities.Mappings{}
	content := entities.Mappings{}
	isVerbatim := false
	for _, value := range patterns {
		if value.Verbatim {
			isVerbatim = true
		}

		// copy path
		l.updateMapping(path, value.Default)
		l.updateMapping(path, value.Path)

		// copy names
		l.updateMapping(name, value.Default)
		l.updateMapping(name, value.Name)

		// copy content
		l.updateMapping(content, value.Default)
		l.updateMapping(content, value.Content)
	}
	if isVerbatim {
		content = entities.Mappings{}
	}
	res := entities.MapResult{
		Path:    l.makeOrderedMap(path),
		Name:    l.makeOrderedMap(name),
		Content: l.makeOrderedMap(content),
	}

	return &res
}

func (l *SombraEngineInteractor) NewFile(path entities.File, paths entities.MapList, names entities.MapList) entities.File {
	dir, fn := filepath.Split(string(path))
	newPath := l.stringProcessor.ProcessString(dir, paths)
	newName := l.stringProcessor.ProcessString(fn, names)
	return entities.File(filepath.Join(newPath, newName))
}

func (l *SombraEngineInteractor) NewContent(content []byte, mappings entities.MapList) []byte {
	return l.stringProcessor.ProcessContent(content, mappings)
}

func (l *SombraEngineInteractor) updateMapping(target entities.Mappings, mapping entities.Mappings) {
	for dk, dv := range mapping {
		target[dk] = dv
	}
}

func (l *SombraEngineInteractor) makeOrderedMap(mappings entities.Mappings) entities.MapList {
	res := make(entities.MapList, 0)

	for key, value := range mappings {
		res = append(res, entities.MapItem{Key: key, Value: value})
	}

	sort.Slice(res, func(item1, item2 int) bool {
		return len(res[item1].Key) > len(res[item2].Key)
	})

	return res
}

func NewSombraEngineInteractor(dirManager DirectoryManagerPort, fileManager FileManagerPort, stringProcessor SombraStringsPort) *SombraEngineInteractor {
	return &SombraEngineInteractor{
		alwaysIgnore: []entities.Wildcard{
			".git/**",
			".idea/**",
			".sombra/**",
			".DS_Store",
			"__pycache__/**/*",
			".mypy_cache/**/*",
		},
		dirManager:      dirManager,
		fileManager:     fileManager,
		stringProcessor: stringProcessor,
	}
}

var _ CliTemplateInitCase = (*CliTemplateInitInteractor)(nil)
