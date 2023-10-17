package main

import "fmt"

func Squaree(m int, ch chan int) {
	square := m * m
	ch <- square // Отправляем результат в канал
}

func main() {
	mas := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(mas)) // Создаем буферизированный канал
	defer close(ch)
	// Закрываем канал
	for _, m := range mas {
		go Squaree(m, ch) // Запускаем горутины для вычисления квадратов
	}

	// Считываем результаты из канала и выводим их
	for range mas {
		square := <-ch // Считываем значение из канала
		fmt.Println("Квадрат=", square)
	}

}
