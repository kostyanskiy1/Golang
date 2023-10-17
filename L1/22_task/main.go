package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	b := new(big.Int)

	// Устанавливаем значение переменных a и b
	a.SetString("2000000", 10) //задать значение > 2^20
	b.SetString("1000000", 10)

	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Println("Умножение:", mul)

	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("Деление:", div)

	add := new(big.Int)
	add.Add(a, b)
	fmt.Println("Сложение:", add)

	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Println("Вычитание:", sub)
}
