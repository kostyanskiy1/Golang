/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

package main

import "fmt"

type Human struct {
	a int
	b string
}

type Action struct {
	Human  //встраиваем структуру Human
	Human1 Human
	abc    int
}

func (h *Human) Show(a string) {
	fmt.Println(h.a, h.b, a)
}

func (a *Action) Show() {
	a.b = "In action"
	//a.Show() так нельязя делать
	fmt.Println(a.a, a.b)
}

func main() {
	var ggg Action

	ggg.a = 1 //присваиваем значения
	ggg.b = "123"
	ggg.Human.a = 5
	ggg.Human1.a = 3

	ggg.Human.Show("пока")
	ggg.Show()
	ggg.Human1.Show("привет")

}
