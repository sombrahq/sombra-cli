package entities

type RepoUpdateInfo struct {
	Branch         string
	CurrentVersion string
}

// File is the relative path of a file in the base directory.
type File string

type FileScanResult struct {
	// File is the relative path without the initial "/"
	File  File
	IsDir bool
	Err   error
}

type MappingType int

type Wildcard string

const (
	MappingDefault MappingType = iota + 1
	MappingPath
	MappingName
	MappingContent
)

type AbstractMappingCandidate struct {
	For      MappingType
	Name     string
	Key      string
	Value    string
	Priority int
}

type FileAnalysis struct {
	Pattern     *Pattern
	IsWildcard  bool
	IsMandatory bool
	Vars        []string
	Exclude     []Wildcard
}

type Mappings map[string]string

type TemplateConfig struct {
	URI     string   `yaml:"uri" validate:"required"`
	Path    string   `yaml:"path,omitempty"`
	Current Version  `yaml:"current,omitempty"`
	Vars    Mappings `yaml:"vars" validate:"required"`
}

type Pattern struct {
	Pattern  Wildcard   `yaml:"pattern" validate:"required"`
	Abstract bool       `yaml:"abstract,omitempty"`
	CopyOnly bool       `yaml:"copy_only,omitempty"`
	Verbatim bool       `yaml:"verbatim,omitempty"`
	Default  Mappings   `yaml:"default,omitempty"`
	Path     Mappings   `yaml:"path,omitempty"`
	Name     Mappings   `yaml:"name,omitempty"`
	Content  Mappings   `yaml:"content,omitempty"`
	Except   []Wildcard `yaml:"except,omitempty"`
}

type TemplateDef struct {
	Vars     []string   `yaml:"vars"`
	Patterns []*Pattern `yaml:"patterns" validate:"required"`
}

type Version string

type MapItem struct {
	Key   string
	Value string
}

type MapList []MapItem

type MapResult struct {
	Path    MapList
	Name    MapList
	Content MapList
}

type SombraTemplateUpdateInfo struct {
	Operation string
	Template  string
	Version   string
}

type SombraUpdateInfo struct {
	Branch  string
	Changes []*SombraTemplateUpdateInfo
}

type SombraDef struct {
	Templates []*TemplateConfig `yaml:"templates" validate:"required"`
}
