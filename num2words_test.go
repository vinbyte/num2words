package num2words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("supported lang", func(t *testing.T) {
		lang := "en"
		assert.NotNil(t, New(lang))
	})
	t.Run("unsupported lang", func(t *testing.T) {
		lang := "fr"
		assert.NotNil(t, New(lang))
	})
}

func TestCustomNegativeWord(t *testing.T) {
	New("en", CustomNegativeWord("negative"))
	assert.Equal(t, "negative", customNegativeWord)
}
