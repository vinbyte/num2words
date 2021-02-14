# Number to Words

[![CircleCI](https://circleci.com/gh/vinbyte/num2words.svg?style=svg)](https://circleci.com/gh/circleci/circleci-docs) [![Coverage Status](https://coveralls.io/repos/github/vinbyte/num2words/badge.svg?branch=main)](https://coveralls.io/github/vinbyte/num2words?branch=main) [![Go Reference](https://pkg.go.dev/badge/github.com/vinbyte/num2words.svg)](https://pkg.go.dev/github.com/vinbyte/num2words)

Golang package to convert number into words. For example : `148` become `one hundred forty-eight`. Inspired by this python library [num2words](https://github.com/savoirfairelinux/num2words).

## Install

Run `go get -u github.com/vinbyte/num2words`

## Usage

```
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
//output : three hundred forty-nine
```

## Test

Run `go test -v`
