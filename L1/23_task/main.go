package main

import "fmt"

func main() {

	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	new := append(a[:i], a[i+1:]...) //Удалить i-ый элемент из слайса
	fmt.Println(new)
}
