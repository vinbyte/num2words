package num2words

import (
	"strconv"
	"strings"
)

func init() {
	l := N2w{
		NumberToWord:       numberToWordEn,
		NumberToOrdinal:    numberToOrdinalEn,
		NumberToOrdinalNum: numberToOrdinalNumEn,
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

func numberToOrdinalEn(number int64) string {
	ordinalWords := make(map[string]string)
	ordinalWords["zero"] = "zeroth"
	ordinalWords["one"] = "first"
	ordinalWords["two"] = "second"
	ordinalWords["three"] = "third"
	ordinalWords["four"] = "fourth"
	ordinalWords["five"] = "fifth"
	ordinalWords["six"] = "sixth"
	ordinalWords["seven"] = "seventh"
	ordinalWords["eight"] = "eighth"
	ordinalWords["nine"] = "ninth"
	ordinalWords["ten"] = "tenth"
	ordinalWords["eleven"] = "eleventh"
	ordinalWords["twelve"] = "twelfth"
	words := numberToWordEn(number)
	splitWord := strings.Split(words, " ")
	lastWord := splitWord[len(splitWord)-1]
	//if any dash separator
	if strings.Contains(lastWord, "-") {
		splitDash := strings.Split(lastWord, "-")
		lastWord = splitDash[len(splitDash)-1]
		splitWord = splitDash
	}
	ords := ordinalWords[lastWord]
	if ords == "" {
		lastLetter := string(lastWord[len(lastWord)-1])
		if lastLetter == "y" {
			lastWord = lastWord[:len(lastWord)-1] + "ie"
		}
		ords = lastWord + "th"
	}
	splitWord[len(splitWord)-1] = ords
	return strings.Join(splitWord, " ")
}

func numberToOrdinalNumEn(number int64) string {
	words := numberToOrdinalEn(number)
	numStr := strconv.FormatInt(number, 10)
	return numStr + words[len(words)-2:]
}
