package num2words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToWordEn(t *testing.T) {
	t.Run("minus", func(t *testing.T) {
		number := -24
		assert.Equal(t, "minus twenty-four", numberToWordEn(int64(number)))
	})
	t.Run("minus custom", func(t *testing.T) {
		number := -24
		customNegativeWord = "negative"
		assert.Equal(t, "negative twenty-four", numberToWordEn(int64(number)))
	})
	t.Run("zero", func(t *testing.T) {
		number := 0
		assert.Equal(t, "zero", numberToWordEn(int64(number)))
	})
	t.Run("hundred", func(t *testing.T) {
		number := 305
		assert.Equal(t, "three hundred five", numberToWordEn(int64(number)))
	})
	t.Run("hundred", func(t *testing.T) {
		number := 101
		assert.Equal(t, "one hundred one", numberToWordEn(int64(number)))
	})
	t.Run("thousand", func(t *testing.T) {
		number := 1315
		assert.Equal(t, "one thousand three hundred fifteen", numberToWordEn(int64(number)))
	})
	t.Run("million", func(t *testing.T) {
		number := 2004355
		assert.Equal(t, "two million four thousand three hundred fifty-five", numberToWordEn(int64(number)))
	})
	t.Run("billion", func(t *testing.T) {
		number := 3050234000
		assert.Equal(t, "three billion fifty million two hundred thirty-four thousand", numberToWordEn(int64(number)))
	})
	t.Run("billion", func(t *testing.T) {
		number := 3050234040
		assert.Equal(t, "three billion fifty million two hundred thirty-four thousand forty", numberToWordEn(int64(number)))
	})
}
