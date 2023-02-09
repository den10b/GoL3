package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func convert(begin string, escapeChar rune) (string, error) {
	builder := strings.Builder{}
	var lastChar rune
	printedEscape := false
	for _, char := range begin {
		if !printedEscape {
			if lastChar == escapeChar {
				if char == escapeChar {
					printedEscape = true
				}
				builder.WriteRune(char)
				lastChar = char
				continue
			}
			if char == escapeChar {
				lastChar = char
				continue
			}

		}
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			builder.WriteRune(char)
			lastChar = char
		} else {
			if lastChar == 0 {
				return "", fmt.Errorf("некорректная строка")
			}
			for j := 1; j < digit; j++ {
				builder.WriteRune(lastChar)
			}
			lastChar = 0
		}
		printedEscape = false

	}
	return builder.String(), nil

}
func main() {
	beginString := "qwe$$4"
	escape := "$"
	runa, _ := utf8.DecodeLastRune([]byte(escape))
	s, err := convert(beginString, runa)
	if err != nil {
		panic(err)
		return
	}
	println(s)
}
