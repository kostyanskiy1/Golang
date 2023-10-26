package main

import (
	"bytes"
	"testing"
)

var addTests = []string{
	"exec test.txt",
	"pwd",
	"ps",
	"echo hello world",
	"fork echo hello world",
	"fork fork ps",
	"kill 2",
	"cd home",
	"pwd",
}

func TestRunCommand(t *testing.T) {
	for _, test := range addTests {
		var output bytes.Buffer

		runCommand(test, &output)

		t.Logf("\nInput:\t%s\nOutput:\t%s", test, output.String())
	}
}
