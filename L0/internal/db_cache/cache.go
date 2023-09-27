package db_cache

import (
	//database "Stan_connection/internal/db"
	"Stan_connection/internal/jsonstr"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

var CC = NewCache()

type Cache struct {
	sync.RWMutex
	cache map[string]jsonstr.ModelJSON
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]jsonstr.ModelJSON),
	}
}

func Get(key string) (jsonstr.ModelJSON, bool) {
	CC.Lock()
	defer CC.Unlock()
	order, found := CC.cache[key] //проверяем в кеше
	return order, found
}

func Set(key string, order jsonstr.ModelJSON) {
	CC.Lock()
	CC.cache[key] = order
	CC.Unlock()
}

func (c *Cache) GetFromDB(db *sqlx.DB) {
	// Запрос к базе данных для получения данных
	models, err := Find(db)
	if err != nil {
		fmt.Printf("Ошибка при загрузке данных из БД: %v", err)
		return
	}

	// Заполнение кэша данными из БД
	for _, m := range models {
		Set(m.OrderUID, m)
	}
	//fmt.Println(c)
	//fmt.Println("Кэш инициализирован из БД")
}
