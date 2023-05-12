package usecase

type Shortener interface {
	Shorten(source string) (string, error)
}
