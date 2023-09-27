package main

import (
	dbcache "Stan_connection/internal/db_cache"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

var db = dbcache.Connect()

func main() {

	defer db.Close()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Обработчик для отображения данных заказа по ID
	http.HandleFunc("/order/", GetOrderByID)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	// Извлечь ID заказа из URL-пути
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Неверный URL", http.StatusBadRequest)
		return
	}

	orderID := parts[2]
	dbcache.CC.GetFromDB(db)
	cachedOrder, found := dbcache.Get(orderID)
	if !found {
		http.Error(w, "Заказ не найден", http.StatusNotFound)
		return
	}

	// Преобразование данных заказа в JSON
	jsonOrder, err := json.MarshalIndent(cachedOrder, "", "\t")
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonOrder))
}
