package main

import (
	"fmt"
	"strconv"
	"strings"
)

func WriteValues(countStr, resultStr *strings.Builder, buffRune *rune) {
	// write "buffRune" "countStr" times in "resultStr" if "buffRune" isn't empty
	if *buffRune != ' ' {
		count, _ := strconv.Atoi(countStr.String())

		if count == 0 {
			count = 1
		}

		for l := 0; l < count; l++ {
			resultStr.WriteRune(*buffRune)
		}
		*buffRune = ' '
		countStr.Reset()
	}
}

func Unpack(str string) string {

	if str == "" {
		return ""
	}

	runeStr := []rune(str)

	var strBuild, countBuff strings.Builder
	buffRune := ' '

	for i := 0; i < len(runeStr); i++ {
		elemStr := string(runeStr[i])
		switch {
		case (elemStr < "0" || elemStr > "9") && elemStr != `\`:
			WriteValues(&countBuff, &strBuild, &buffRune)
			buffRune = runeStr[i]
		case "0" < elemStr && elemStr <= "9" && buffRune != ' ':
			countBuff.WriteRune(runeStr[i])
		case elemStr == `\` && i+1 < len(runeStr):
			WriteValues(&countBuff, &strBuild, &buffRune)
			i++
			buffRune = runeStr[i]
		}
	}

	WriteValues(&countBuff, &strBuild, &buffRune)
	if strBuild.Len() == 0 {
		return "(некорректная строка)"
	}

	return strBuild.String()
}

func main() {
	str := `ab9cd`
	fmt.Println("Распакованная строка: ", Unpack(str))
}
