# Number to Word

Golang package to convert number into word. For example : `48` become `forty-eight`. Inspired by this python library [num2words](https://github.com/savoirfairelinux/num2words).

## Install

`go get -u github.com/vinbyte/num2words`

## Usage

```
package main

import (
	"fmt"
	"github/vinbyte/num2words"
)

func main() {
	n2w := num2words.New("en")
	res := n2w.NumberToWord(349)
	fmt.Println(res)
}
//output : three hundred forty-nine
```

## Test

`go test -v`