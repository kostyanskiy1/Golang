package main

import (
	"testing"
)

type addTest struct {
	host string
	err  error
}

var addTests = []addTest{
	addTest{"0.beevik-ntp.pool.ntp.org", nil},
}

func TestGetTime(t *testing.T) {
	for _, test := range addTests {
		if _, output := GetTime(test.host); output != test.err {
			t.Errorf("Output %q not equal to expected %q", output, test.err)
		}
	}
}
