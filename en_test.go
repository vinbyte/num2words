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

func TestNumberToOrdinalEn(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		number := int64(0)
		assert.Equal(t, "zeroth", numberToOrdinalEn(number))
	})
	t.Run("1", func(t *testing.T) {
		number := int64(1)
		assert.Equal(t, "first", numberToOrdinalEn(number))
	})
	t.Run("2", func(t *testing.T) {
		number := int64(2)
		assert.Equal(t, "second", numberToOrdinalEn(number))
	})
	t.Run("11", func(t *testing.T) {
		number := int64(11)
		assert.Equal(t, "eleventh", numberToOrdinalEn(number))
	})
	t.Run("15", func(t *testing.T) {
		number := int64(15)
		assert.Equal(t, "fifteenth", numberToOrdinalEn(number))
	})
	t.Run("21", func(t *testing.T) {
		number := int64(21)
		assert.Equal(t, "twenty first", numberToOrdinalEn(number))
	})
	t.Run("22", func(t *testing.T) {
		number := int64(22)
		assert.Equal(t, "twenty second", numberToOrdinalEn(number))
	})
	t.Run("130", func(t *testing.T) {
		number := int64(130)
		assert.Equal(t, "one hundred thirtieth", numberToOrdinalEn(number))
	})
}

func TestNumberToOrdinalNumEn(t *testing.T) {
	t.Run("small", func(t *testing.T) {
		number := int64(48)
		assert.Equal(t, "48th", numberToOrdinalNumEn(number))
	})
	t.Run("big", func(t *testing.T) {
		number := int64(3000000000000)
		assert.Equal(t, "3000000000000th", numberToOrdinalNumEn(number))
	})
}
