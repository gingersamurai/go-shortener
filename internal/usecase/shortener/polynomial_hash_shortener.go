package shortener

import (
	"errors"
	"math"
	"strings"
	"unicode/utf8"
)

var (
	ErrLengthMustBePositive = errors.New("length must be positive integer")
	ErrGotMaxCount          = errors.New("got max count")
)

type PolynomialHashShortener struct {
	length   int
	count    int
	maxCount int
	alphabet string
}

func NewPolynomialHashShortener(length int) (*PolynomialHashShortener, error) {
	if length <= 0 {
		return nil, ErrLengthMustBePositive
	}

	var (
		lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
		uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits           = "0123456789"
		extra            = "_"
	)
	alphabet := strings.Builder{}
	alphabet.WriteString(lowercaseLetters)
	alphabet.WriteString(uppercaseLetters)
	alphabet.WriteString(digits)
	alphabet.WriteString(extra)

	maxCount := int(math.Pow(float64(utf8.RuneCountInString(alphabet.String())), float64(length)))

	return &PolynomialHashShortener{
		length:   length,
		maxCount: maxCount,
		alphabet: alphabet.String(),
	}, nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (phs *PolynomialHashShortener) Shorten(string) (string, error) {
	if phs.count == phs.maxCount {
		return "", ErrGotMaxCount
	}
	result := strings.Builder{}
	curCount := phs.count
	for i := 0; i < phs.length; i++ {
		pos := curCount % utf8.RuneCountInString(phs.alphabet)
		result.WriteRune([]rune(phs.alphabet)[pos])
		curCount /= utf8.RuneCountInString(phs.alphabet)
	}
	phs.count++
	return reverse(result.String()), nil
}
