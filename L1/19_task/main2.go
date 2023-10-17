package main

import (
	"fmt"
)

func main() {
	old := "главрыба"

	runes := []rune(old)                                  //переводим в руны
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { //меняем месатами с начала и конца
		runes[i], runes[j] = runes[j], runes[i]
	}
	new := string(runes)

	fmt.Println(new)
}
