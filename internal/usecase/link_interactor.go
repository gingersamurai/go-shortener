package usecase

type LinkInteractor struct {
	storage Storage
}

func (li *LinkInteractor) AddLink(source string) (string, error) {
	return "", nil
}

func (li *LinkInteractor) GetLink(mappedLink string) (string, error) {
	return "", nil
}
