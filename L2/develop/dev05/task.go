package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Cmd struct {
	After      uint
	Before     uint
	Context    uint
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func (cmd *Cmd) Usage() {
	fmt.Printf("Usage of %s:\ngo run task.go [-A|-B|-C] [-c] [-i] [-v] [-F] [-n] [pattern] [input_file.txt] \n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

func (cmd *Cmd) Parse() {
	flag.Usage = cmd.Usage
	flag.UintVar(&cmd.After, "A", 0, "печатать +N строк после совпадения")
	flag.UintVar(&cmd.Before, "B", 0, "печатать +N строк до совпадения")
	flag.UintVar(&cmd.Context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&cmd.Count, "c", false, "количество строк")
	flag.BoolVar(&cmd.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&cmd.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&cmd.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&cmd.LineNum, "n", false, "печатать номера строки")
	flag.Parse()
}

func main() {
	var cmd = Cmd{}
	cmd.Parse()

	pattern := flag.Args()[0]
	var numberOfString []int

	if cmd.Fixed {
		pattern = `\Q` + pattern + `\E`
	}

	if cmd.IgnoreCase {
		pattern = `(?i)` + pattern
	}

	reg := regexp.MustCompile(pattern)

	// ------------ READ DATA FROM FILE ------------
	dataFromFile, err := os.ReadFile(flag.Args()[1])
	check(err)

	// ------------ PROCESSING DATA FROM FILE ------------

	lines := strings.Split(string(dataFromFile), "\n")
	for i, v := range lines {
		if reg.Match([]byte(v)) {
			numberOfString = append(numberOfString, i)
		}
	}

	var resultOutput []string
	before := int(math.Max(float64(cmd.Context), float64(cmd.Before)))
	after := int(math.Max(float64(cmd.Context), float64(cmd.After)))

	if cmd.Invert {
		ind := 0
		for i, v := range lines {
			if i != numberOfString[ind] {
				resultOutput = append(resultOutput, v)
			} else {
				if ind+1 < len(numberOfString) {
					ind++
				}
				continue
			}
		}
	} else {
		for _, v := range numberOfString {
			if v-before >= 0 {
				resultOutput = append(resultOutput, lines[v-before:v]...)
			} else {
				resultOutput = append(resultOutput, lines[:v]...)
			}
			resultOutput = append(resultOutput, lines[v])
			if v+after+1 > len(lines) {
				resultOutput = append(resultOutput, lines[v+1:]...)
			} else {
				resultOutput = append(resultOutput, lines[v+1:v+after+1]...)
			}
		}
	}

	if cmd.Count {
		resultOutput = append(resultOutput, fmt.Sprintf("Count of lines: %v", len(numberOfString)))
	}

	if cmd.LineNum {
		var sb strings.Builder
		sb.WriteString("Numbers of line with match: ")
		for _, v := range numberOfString {
			sb.WriteString(strconv.Itoa(v))
			sb.WriteString(" ")
		}
		resultOutput = append(resultOutput, sb.String())
	}

	for _, v := range resultOutput {
		fmt.Println(v)
	}
}
