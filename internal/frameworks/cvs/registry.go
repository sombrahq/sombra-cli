package cvs

import (
	"github.com/sombrahq/sombra-cli/internal/core/usecases"
	"github.com/sombrahq/sombra-cli/internal/frameworks/cvs/git"
	"github.com/sombrahq/sombra-cli/internal/frameworks/logger"
)

type RegistryType map[string]usecases.RepositoryFactory

var registry = RegistryType{
	"git": git.Factory,
}

func For(url string) (usecases.RepositoryPort, error) {
	logger.Info("Calling LocalRepoFactory for URL: " + url)
	factory := registry["git"]
	repoPort, err := factory(url)
	if err != nil {
		logger.Error("Failed to get LocalRepoPort", err)
		return nil, err
	}
	logger.Info("Successfully obtained LocalRepoPort for URL: " + url)
	return repoPort, nil
}
