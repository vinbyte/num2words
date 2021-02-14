package num2words

import (
	"strings"
)

func init() {
	l := N2w{
		NumberToWord: numberToWordEn,
	}
	languages["en"] = &l
}

func numberToWordEn(number int64) string {
	megaWords := []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"}
	unitWords := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tenWords := []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	teenWords := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	var words []string
	if number < 0 {
		minus := "minus"
		if customNegativeWord != "" {
			minus = customNegativeWord
		}
		words = append(words, minus)
		number *= -1
	}
	threeDigits := splitByThousand(number)
	if len(threeDigits) == 0 {
		return "zero"
	}
	level := len(threeDigits) - 1
	for _, num := range threeDigits {
		if num == 0 {
			level--
			continue
		}
		hundred := num / 100 % 10
		ten := num / 10 % 10
		unit := num % 10
		if hundred > 0 {
			words = append(words, unitWords[hundred], "hundred")
		}
		if ten == 0 {
			words = append(words, unitWords[unit])
		} else if ten == 1 {
			words = append(words, teenWords[unit])
		} else {
			if unit > 0 {
				word := tenWords[ten] + "-" + unitWords[unit]
				words = append(words, word)
			} else {
				words = append(words, tenWords[ten])
			}
		}
		if mega := megaWords[level]; mega != "" {
			words = append(words, mega)
		}
		level--
	}

	return strings.Join(words, " ")
}
