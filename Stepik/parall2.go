package main

/*В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.

Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу. Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая управление.*/

import (
	"fmt"
)

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {

	channel := make(chan int)

	go func() {
		defer close(channel)
		sum := 0
		for {
			select { // Оператор select
			case x := <-arguments:
				sum += x

			case <-done:
				channel <- sum
				return

			}
		}
	}()
	return channel
}

func main() {
	ch1 := make(chan int)
	stop := make(chan struct{})
	r := calculator(ch1, stop)
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5
	close(stop)
	fmt.Println(<-r)
}
