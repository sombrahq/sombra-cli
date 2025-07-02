package usecases

type RepositoryPort interface {
	Clone() error
	Clean() error
	Dir() string
	Use(version string) (string, error)
	Diff(commit string) ([]byte, error)
	GetTags() ([]string, error)
}

type PatchPort interface {
	Apply(dir string, patch []byte) error
}

type RepositoryFactory func(uri string) (RepositoryPort, error)
