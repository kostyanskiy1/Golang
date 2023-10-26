package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и go lint.
*/

type Cmd struct {
	ChooseColumnForSort     uint
	SortByNumericValue      bool
	SortInReverseOrder      bool
	DontWriteNonUniqueLines bool
	SortStyle               func([]string)
}

func (cmd *Cmd) Usage() {
	fmt.Printf("Usage of %s:\ngo run task.go [-k num_of_column] [-n|-r] [-u] [input_file.txt] [output_file.txt] \n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

func (cmd *Cmd) Parse() {
	flag.Usage = cmd.Usage
	flag.UintVar(&cmd.ChooseColumnForSort, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&cmd.SortByNumericValue, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&cmd.SortInReverseOrder, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&cmd.DontWriteNonUniqueLines, "u", false, "не выводить повторяющиеся строки")
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

	// READ DATA FROM FILE
	dataFromFile, err := os.ReadFile("dev03/input.txt")
	check(err)

	// PROCESSING DATA FROM FILE

	lines := strings.Split(string(dataFromFile), "\n")

	// FLAG HANDLER
	if cmd.SortByNumericValue {
		cmd.SortStyle = SortByNumericValue
	} else if cmd.SortInReverseOrder {
		cmd.SortStyle = SortInReverseOrder
	}

	if cmd.ChooseColumnForSort != 0 {
		strMatrix := CreateMatrixOfStringsFromData(lines, cmd.DontWriteNonUniqueLines, cmd.ChooseColumnForSort)
		cmd.SortStyle(strMatrix[cmd.ChooseColumnForSort])
		lines = MatrixToLines(strMatrix)
	} else if cmd.SortByNumericValue || cmd.SortInReverseOrder {
		cmd.SortStyle(lines)
	}

	// WRITE TO FILE
	// create output file
	f, err := os.Create("dev03/out.txt")
	check(err)
	WriteLineToFile(f, lines, cmd.DontWriteNonUniqueLines)
}

func SortByNumericValue(strArr []string) {
	sort.Strings(strArr)
}

func SortInReverseOrder(strArr []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(strArr)))
}

func CreateMatrixOfStringsFromData(lines []string, repeatingStringsFlag bool, numerOfCoulumn uint) [][]string {
	var set = make(map[string]struct{})

	// split input file by lines

	var strMatrixLines = make([][]string, len(lines))

	countOfColumn := 0
	// split input lines by spaces (we got slice of slices where main slice is line, but we need column as main slice)
	for i, val := range lines {
		strMatrixLines[i] = strings.Split(val, " ")
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
			if _, ok := set[strMatrixLines[j][i]]; ok && repeatingStringsFlag && numerOfCoulumn == uint(i) {
				strMatrixColumn[i] = append(strMatrixColumn[i], "")
			} else {
				if numerOfCoulumn == uint(i) {
					set[strMatrixLines[j][i]] = struct{}{}
				}
				strMatrixColumn[i] = append(strMatrixColumn[i], strMatrixLines[j][i])
			}
		}
	}
	return strMatrixColumn
}

func MatrixToLines(matrix [][]string) []string {
	var buffSt strings.Builder
	var buffStArr []string
	for i := range matrix[0] {
		for j := range matrix {
			buffSt.WriteString(matrix[j][i])
			buffSt.WriteString(" ")
		}
		buffStArr = append(buffStArr, buffSt.String())
		buffSt.Reset()
	}
	return buffStArr
}

func WriteLineToFile(f *os.File, line []string, repeatingStringsFlag bool) {
	var set = make(map[string]struct{})
	for i := range line {
		if _, ok := set[line[i]]; ok && repeatingStringsFlag {

		} else {
			_, err := f.Write([]byte(line[i] + "\n"))
			check(err)
			set[line[i]] = struct{}{}
		}
	}
}
