package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct { //структура-счетчик
	count int
	atom  int32
	mx    sync.Mutex
}

func (c *Counter) Inc() { //инкремент в конкурентной среде
	c.mx.Lock()
	c.count++ //с помощью мьютекса
	c.mx.Unlock()
	atomic.AddInt32(&c.atom, 1) //с помощью Atomic
}

func main() {
	wg := sync.WaitGroup{}
	var c Counter
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("C.count=", c.count, "C.atom=", c.atom)
}
