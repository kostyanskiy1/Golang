package main

import "testing"

type addTest struct {
	input, expected string
}

var addTests = []addTest{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"45", "(некорректная строка)"},
	{"", ""},
	{`qwe\4\5`, "qwe45"},
	{`qwe\45 `, "qwe44444"},
	{`qwe\\5`, "qwe\\\\\\\\\\"},
	{"5d", "d"},
}

func TestUnpackString(t *testing.T) {
	for _, test := range addTests {
		if output := Unpack(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
