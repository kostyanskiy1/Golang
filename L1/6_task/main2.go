package main

import (
	"fmt"
)

func main() {
	done := make(chan string)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		go myFunc(i, done)
	}

	for a := range done {

		fmt.Println(a)
	}

}

func myFunc(i int, done chan string) { //мы ждем, что канал done вернет нам какое-то значение
	fmt.Println("hello from ", i)
	done <- fmt.Sprint("closed hello from ", i)
	close(done) //закрывает этот канал после завершения работы
}
