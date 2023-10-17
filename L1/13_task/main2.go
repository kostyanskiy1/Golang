package main

import "fmt"

func main() {
	a := 3
	b := 5
	fmt.Println("До перестановки:", a, b)
	a, b = b, a
	fmt.Println("После перестановки:", a, b)

}
