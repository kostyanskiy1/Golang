package main

import (
	"fmt"
	"sync"
)

func do(ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		switch v := i.(type) { //определить тип переменной из переменной типа interface{}
		case int:
			fmt.Println("Получили int=", v)
		case string:
			fmt.Println("Получили string=", v)
		case bool:
			fmt.Println("Получили bool=", v)
		case chan int:
			fmt.Println("Получили chan int=", v)
		default:
			fmt.Printf("Я не знаю такого типа %T!\n", v)
		}

	}
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan interface{}) //канал переменной типа interface{}.
	cha := make(chan int)

	wg.Add(1)
	go do(ch, &wg)

	ch <- 5 //передаем  int, string, bool, channel
	ch <- true
	ch <- "str"
	ch <- cha

	close(ch)
	wg.Wait()
}
