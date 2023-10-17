package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := new(sync.WaitGroup)

	mu := new(sync.Mutex)
	mas := [5]int{2, 4, 6, 8, 10}
	var sum int
	var sum2 int64
	for _, m := range mas {
		wg.Add(2) // Увеличиваем счетчик горутин в группе
		go sq(m, wg, &sum, mu)
		go sq2(m, &sum2, wg) // Вызываем функцию в отдельной горутине
	}

	wg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("sum=", sum, "sum2=", sum2)
}

func sq(m int, wg *sync.WaitGroup, sum *int, mu *sync.Mutex) {
	defer wg.Done() // уменьшаем счетчик горутин в группе
	mu.Lock()       // Заблокировать мьютекс перед изменением общей суммы
	*sum += m * m
	mu.Unlock()

}

func sq2(m int, sum2 *int64, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик горутин в группе
	atomic.AddInt64(sum2, int64(m*m))
}
