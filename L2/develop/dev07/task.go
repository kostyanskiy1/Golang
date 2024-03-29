package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===
Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	for _, c := range channels {
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			close(out)
		}(c)
	}
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			c <- []int{1, 4, 5}
		}()
		return c
	}
	start := time.Now()
	a := <-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(2*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Println("a=", a)
	fmt.Printf("done after %v", time.Since(start))
}
