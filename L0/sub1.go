package main

import (
	dbcache "Stan_connection/internal/db_cache"
	"Stan_connection/internal/jsonstr"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("prod", "sub-1")
	defer sc.Close()

	db := dbcache.Connect()
	defer db.Close()

	msgCh := make(chan jsonstr.ModelJSON)
	var model jsonstr.ModelJSON

	subscription, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))

		dbcache.GetFromStan(m, model, msgCh, dbcache.CC) //берем из сообщения и отправляем в канал и кэш

	}, stan.DeliverAllAvailable())
	if err != nil {
		log.Fatalf("Ошибка подписки на канал: %v", err)
	}
	defer subscription.Close()

	dbcache.Push(msgCh, db) //добавляем в бд
	fmt.Println("Запись отправлена в бд")

}
