/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

func Square(m int, wg *sync.WaitGroup) {
	square := m * m
	fmt.Println(m, square)
	defer wg.Done() // Уменьшаем счетчик, когда функция завершает выполнение
}

func main() {
	mas := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // Создаем объект WaitGroup для синхронизации горутин

	for _, m := range mas {
		wg.Add(1)         // Увеличиваем счетчик для каждой горутины
		go Square(m, &wg) // Запускаем горутину
	}
	wg.Wait() // Ожидаем завершения всех горутин

}
