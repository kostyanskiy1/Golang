package main

import (
	"fmt"
	"strings"
)

func main() {

	old := "snow dog sun"
	new := reverse(old)

	fmt.Println(new)
}

func reverse(old string) string { //переворачивает слова в строке.
	a := strings.Split(old, " ") //разделяем в срез строк
	b := ""
	for i := len(a) - 1; i >= 0; i-- { //передаем в вывод в обратном порядке
		b += a[i] + " "
	}
	return b
}
