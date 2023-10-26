package main

import (
	"fmt"
	"strings"
	"testing"
)

type addTest struct {
	input    []string
	expected string
}

var addTests = []addTest{
	{[]string{"тяпка", "пятак", "лиСток", "пЯтка", "слиТок", "столик"}, "пятак: &[пятак пятка тяпка]листок: &[листок слиток столик]"},
}

func TestFindAnogram(t *testing.T) {
	for _, test := range addTests {
		var output strings.Builder
		set := FindAnogram(&test.input)
		for k, v := range *set {
			output.WriteString(fmt.Sprintf("%s: %v", k, v))
		}
		if output.String() != test.expected {
			t.Errorf("Output %s not equal to expected %s", output.String(), test.expected)
		}
	}
}
