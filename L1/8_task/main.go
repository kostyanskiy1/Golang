package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int64 = 352537456568575757
	v := strconv.FormatInt(a, 2) //переводим в двоичную систему
	fmt.Println(v)
	fmt.Println("a=", a)

	num := 3

	//w |= 1 << 12 Чтобы установить 12-й бит слова w на единицу
	//Чтобы очистить (установить на ноль) 12-й бит слова w: w &^= 1<<12
	a ^= 1 << num //меняем знак по номеру на противоположный
	fmt.Printf("res=%064b\n", a)
	fmt.Println("a=", a)

}
