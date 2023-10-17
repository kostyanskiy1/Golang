package main

import (
	"fmt"
	"sort"
)

func main() {

	mas := []int{1, 5, 7, 9}
	target := 5
	num := sort.SearchInts(mas, target) //бинарный поиск

	if mas[num] == target {
		fmt.Printf("Элемент %d на %d месте", target, num+1)
	} else {
		fmt.Println("Не найдено")
	}

}
