package main

import (
	"fmt"
	"time"
)

func sleepyGopher(c chan int) {
	for i := 0; ; i++ {
		fmt.Printf("gopher %d sleep\n", i)
		time.Sleep(time.Millisecond * 200)
		c <- i
	}
}

func main() {
	var n int
	fmt.Println("Введите время:")
	fmt.Scan(&n)
	timeout := time.After(time.Duration(n) * time.Second)

	c := make(chan int, 5)

	go sleepyGopher(c)

	for {
		select { // Оператор select
		case gopherID := <-c: // Ждет, когда проснется гофер
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout: // Ждет окончания времени
			fmt.Println("Время вышло")

			return // Сдается и возвращается
		}
	}
}
