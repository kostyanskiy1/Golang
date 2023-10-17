package main

import (
	"fmt"
	"sync"
)

// Функция для чтения чисел из in, удвоения их и записи в out
func x2(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range in { //считываем, умножаем и отправляем
		result := x * 2
		out <- result
	}
	close(out) //по завершению закрываем канал для отправки
}

func res(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range out { //  вывод
		fmt.Println(result)
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go x2(in, out, &wg) //заупскаем горутины
	go res(out, &wg)

	mas := []int{1, 2, 3, 4, 5}

	for _, m := range mas { // Пишем числа в первый канал
		in <- m
	}
	close(in) // Закрываем первый канал, чтобы завершить горутину чтения
	wg.Wait()
}
