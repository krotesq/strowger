package source

type service struct {
	repository *repository
}

func newService(repository *repository) *service {
	return &service{repository: repository}
}
