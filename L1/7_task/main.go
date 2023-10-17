package main

import (
	"fmt"
	"sync"
)

func main() {

	myMap := make(map[string]int)

	var mutex sync.RWMutex
	num := 5

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup
	wg.Add(num)

	// Горутины для записи в map
	for i := 0; i < num; i++ {
		go func(n int) {
			mutex.Lock() // блoкируем мьютекс
			defer mutex.Unlock()

			myMap[fmt.Sprintf("key%d", n)] = n // Добавляем элемент в map

			fmt.Printf("Written: key%d\n", n)
			wg.Done() // Уменьшаем счетчик WaitGroup
		}(i)
	}
	wg.Wait() // Ожидаем завершения всех горутин
	wg.Add(num)
	// Горутины для чтения из map
	for i := 0; i < num; i++ {
		go func(n int) {
			// Захватываем мьютекс
			mutex.RLock()
			defer mutex.RUnlock()
			// Читаем значение из map
			val := myMap[fmt.Sprintf("key%d", n)]

			fmt.Printf("Read: key%d = %d\n", n, val)

			// Уменьшаем счетчик WaitGroup
			wg.Done()
		}(i)
	}

	wg.Wait() // Ожидаем завершения всех горутин
}
