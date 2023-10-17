package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(i int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		// Воркеры читают данные из канала и выводят их в stdout.
		fmt.Printf("Воркер %d: %s\n", i, data)
	}
	fmt.Printf("Воркер %d завершил работу.\n", i)
}

func main() {
	ch := make(chan string)

	// Создаем канал для ожидания завершения работы всех воркеров.
	var wg sync.WaitGroup

	var n int
	fmt.Println("Введите количетсво воркеров:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}
	// Создаем канал для приема сигналов
	sig := make(chan os.Signal, 1)

	// Говорим пакету signal пересылать сигналы SIGINT и SIGTERM в канал
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем горутину, которая будет ждать сигнала
	go func() {
		<-sig
		// При получении сигнала выводим сообщение и завершаем программу
		fmt.Println("Получен сигнал завершения")
		os.Exit(0)
	}()

	// Ваш код программы...

	// Программа будет ждать на этой точке, пока не будет получен сигнал завершения
	//<-sig

	for i := 1; i <= 30; i++ {
		ch <- fmt.Sprintf("Сообщение %d", i)
	}

	wg.Wait()
	fmt.Println("Все воркеры завершили работу.")
}
