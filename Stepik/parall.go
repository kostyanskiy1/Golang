package main

/*Необходимо написать функцию func merge2Channels , она должна n раз сделать следующее

прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
вычислить f(x1) + f(x2)
записать полученное значение в out
Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 2

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	wg := new(sync.WaitGroup)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			// Запускаем n экземпляров горутины
			wg.Add(1)

			x1 := <-in1
			x2 := <-in2

			res := fn(x1) + fn(x2)

			out <- res
		}
	}(wg)

	wg.Wait()

}
