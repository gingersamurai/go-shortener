package usecase

type Shortener interface {
	ShortenLink(source string) (string, error)
}
