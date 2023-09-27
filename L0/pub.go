package main

import (
	"Stan_connection/internal/jsonstr"
	"encoding/json"
	"fmt"
	"os"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("prod", "simple-pub")
	defer sc.Close()

	for {
		fmt.Println("Введите имя файла JSON:")
		var fl string
		fmt.Scan(&fl)
		dataFromFile, err := os.ReadFile(fl)
		if err != nil {
			fmt.Println("Ошибка с чтением:", err)
		}

		var js jsonstr.ModelJSON
		if err := json.Unmarshal(dataFromFile, &js); err != nil { //тоже как валидация
			fmt.Println(err)
			continue
		}

		strjs, err := json.Marshal(js)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Simple Synchronous Publisher
		sc.Publish("foo", []byte(strjs)) // does not return until an ack has been received from NATS Streaming
		fmt.Println("Отправлено сообщение")
	}
}
