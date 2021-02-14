package main

import (
	"fmt"

	"github.com/vinbyte/num2words"
)

func main() {
	n2w := num2words.New("en")
	res := n2w.NumberToWord(349)
	fmt.Println(res)
}
