package main

import (
	"fmt"
	"time"
)

func main() {

	mas := [5]int{2, 4, 6, 8, 10}

	ch := make(chan int)

	defer close(ch)
	square2(ch)
	for _, m := range mas {
		ch <- m //отправляем по каналу значение
	}

	time.Sleep(time.Millisecond * 300) //ожидем завершения горутин
}

func square2(ch chan int) {
	var sum int

	go func() { //

		for i := 0; i < 5; i++ {
			select { // Оператор select
			case a := <-ch: // Ждет, когда проснется гофер
				fmt.Println("a=", a, "sq=", a*a)
				sum += a * a

			}
		}
		fmt.Println("Сумма sq=", sum)
	}()

}
