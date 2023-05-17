package polynomial_hash_shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolynomialHashShortener(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		phs, err := NewPolynomialHashShortener(4)
		assert.NoError(t, err)

		gotResult, err := phs.Shorten("extra string")
		assert.NoError(t, err)
		assert.Equal(t, gotResult, "aaaa")

		gotResult, err = phs.Shorten("extra string")
		assert.NoError(t, err)
		assert.Equal(t, gotResult, "aaab")
	})

	t.Run("bad length", func(t *testing.T) {
		_, err := NewPolynomialHashShortener(-4)
		assert.ErrorIs(t, err, ErrLengthMustBePositive)
	})

	t.Run("got max count", func(t *testing.T) {
		phs, err := NewPolynomialHashShortener(1)
		assert.NoError(t, err)
		for i := 0; i < 63; i++ {
			_, err := phs.Shorten("extra string")
			assert.NoError(t, err)
		}

		_, err = phs.Shorten("extra string")
		assert.ErrorIs(t, err, ErrGotMaxCount)
	})
}
