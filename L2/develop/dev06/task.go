package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и go lint.
*/

type Cmd struct {
	ChooseField   uint
	Separator     string
	WithSeparator bool
}

func (cmd *Cmd) Usage() {
	fmt.Printf("Usage of %s:\ngo run task.go [-f [numer_of_column]] [-d [separator]] [-s]  [input_file.txt] \n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

func (cmd *Cmd) Parse() {
	flag.Usage = cmd.Usage
	flag.UintVar(&cmd.ChooseField, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&cmd.Separator, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&cmd.WithSeparator, "s", false, "только строки с разделителем")
	flag.Parse()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var cmd = Cmd{}
	cmd.Parse()

	// ------------ READ DATA FROM FILE ------------
	dataFromFile, err := os.ReadFile(flag.Args()[0])
	check(err)

	// ------------ PROCESSING DATA FROM FILE ------------

	lines := strings.Split(string(dataFromFile), "\n")

	if cmd.ChooseField != 0 {
		strMatrixColumn := CreateMatrixOfStringsFromData(lines, cmd.Separator)
		fmt.Println(strings.Join(strMatrixColumn[cmd.ChooseField-1], "\n"))
	} else if cmd.WithSeparator {
		for _, v := range lines {
			if strings.Contains(v, cmd.Separator) {
				fmt.Println(v)
			}
		}
	} else {
		fmt.Println(strings.Join(lines, "\n"))
	}
}

func CreateMatrixOfStringsFromData(lines []string, separator string) [][]string {
	// split input file by lines

	var strMatrixLines = make([][]string, len(lines))

	countOfColumn := 0
	// split input lines by spaces (we got slice of slices where main slice is line, but we need column as main slice)
	for i, val := range lines {
		strMatrixLines[i] = strings.Split(val, separator)
		if len(strMatrixLines[i]) > countOfColumn {
			countOfColumn = len(strMatrixLines[i])
		}
	}

	// format matrix (add empty values in lines for formatting matrix)
	for i := range strMatrixLines {
		for len(strMatrixLines[i]) < countOfColumn {
			strMatrixLines[i] = append(strMatrixLines[i], "")
		}
	}

	// create matrix of strings which have column in main slice
	var strMatrixColumn = make([][]string, countOfColumn)
	for i := range strMatrixColumn {
		for j := range strMatrixLines {
			strMatrixColumn[i] = append(strMatrixColumn[i], strMatrixLines[j][i])
		}
	}
	return strMatrixColumn
}
