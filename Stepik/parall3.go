package main

/*
Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.

в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.*/

import (
	"fmt"
)

func calculator2(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {

	channel := make(chan int)

	go func() {
		defer close(channel)
		for {
			select { // Оператор select

			case x := <-firstChan: // Ждет, когда проснется гофер
				channel <- x * x

			case x := <-secondChan: // Ждет окончания времени
				channel <- x * 3

			case <-stopChan:
				return
			}
		}
	}()

	return channel
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator2(ch1, ch2, stop)

	ch1 <- 2
	fmt.Println(<-r)
	ch2 <- 2
	fmt.Println(<-r)
	close(stop)

}
