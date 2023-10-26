package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// требуется только ниже для обработки примера

func main() {

	fmt.Println("Запуск сервера")
	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("Ошибка в прослушивании порта:", ln)
	}

	// Открываем порт
	conn, _ := ln.Accept()
	defer conn.Close()

	// Запускаем цикл
	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Println()
		fmt.Print("Получено сообщение: ", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
