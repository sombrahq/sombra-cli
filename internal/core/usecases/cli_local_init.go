package usecases

type CliLocalInitCase interface {
	DoLocalInit(target, uri string) error
}

type CliLocalInitInteractor struct {
	localInitCase LocalInitCase
}

func (l *CliLocalInitInteractor) DoLocalInit(target, uri string) error {
	return l.localInitCase.LocalInit(target, uri)
}

func NewCliLocalInitInteractor(localInitCase LocalInitCase) *CliLocalInitInteractor {
	return &CliLocalInitInteractor{localInitCase: localInitCase}
}

var _ CliLocalInitCase = (*CliLocalInitInteractor)(nil)
