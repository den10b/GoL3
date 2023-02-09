package main

import (
	"testing"
	"unicode/utf8"
)

type convTest struct {
	begin, escape, expected string
}

var convTests = []convTest{
	{"a4bc2d5e", "$", "aaaabccddddde"},
	{"abcd", "$", "abcd"},
	{"", "$", ""},
	{"qwe$$4", "$", "qwe$$$$"},
	{"qwe*4*5", "*", "qwe45"},
	{"qwe$45", "$", "qwe44444"},
}

func TestConvert(t *testing.T) {
	for _, test := range convTests {
		runa, _ := utf8.DecodeLastRune([]byte(test.escape))
		if output, _ := convert(test.begin, runa); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}

}
