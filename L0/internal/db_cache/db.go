package db_cache

import (
	"Stan_connection/internal/jsonstr"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
)

func Connect() *sqlx.DB {
	var db *sqlx.DB
	connStr := "user=postgres password=12345 dbname=postgres sslmode=disable"
	var err error
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка при соединении c БД", err)
	} else {
		fmt.Println("Postgre Connected!")
	}
	return db
}

func GetFromStan(m *stan.Msg, model jsonstr.ModelJSON, msgCh chan jsonstr.ModelJSON, CC *Cache) {

	if err := json.Unmarshal(m.Data, &model); err != nil {

		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type) //валидация
			retError := fmt.Errorf("Invalid json field[%s] type: got: %s, want: %v", e.Field, e.Value, e.Type)
			fmt.Println(retError)

		case *json.InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)

		default:
			fmt.Println(err)
		}
	} else {
		// Кэшировать полученные данные
		cachedOrder, found := Get(model.OrderUID)
		if found {
			fmt.Printf("Данные уже есть в кэше: %+v\n", cachedOrder)
		} else {
			msgCh <- model
			Set(model.OrderUID, model)
		}
	}
}

func Push(msgCh chan jsonstr.ModelJSON, tx *sqlx.DB) {
	for msg := range msgCh {
		// Прочитайте данные сообщения
		data := msg
		textjson, err := json.Marshal(data.Items)
		if err != nil {
			fmt.Println(err)
		}

		result1, err := tx.Exec("insert into Paymentt values ($1, $2,$3, $4,$5, $6,$7, $8,$9, $10)",
			data.Payment.Transaction, data.Payment.RequestID, data.Payment.Currency, data.Payment.Provider, data.Payment.Amount, data.Payment.PaymentDt,
			data.Payment.Bank, data.Payment.DeliveryCost, data.Payment.GoodsTotal, data.Payment.CustomFee)
		if err != nil {
			fmt.Println("Ошибка отправки в БД:", err)
		}

		result2, err := tx.Exec("insert into Deliveryt values ($1, $2,$3, $4,$5, $6,$7)", data.Delivery.Name,
			data.Delivery.Phone, data.Delivery.Zip, data.Delivery.City, data.Delivery.Address, data.Delivery.Region, data.Delivery.Email)
		if err != nil {
			fmt.Println("Ошибка отправки в БД:", err)
		}

		result, err := tx.Exec("insert into wildberry values ($1, $2,$3, $4,$5, $6,$7, $8,$9, $10, $11, $12, $13, $14)",
			data.OrderUID, data.TrackNumber, data.Entry, data.Delivery.Phone, data.Payment.Transaction, string(textjson),
			data.Locale, data.InternalSignature, data.CustomerID, data.DeliveryService, data.Shardkey, data.SmID, data.DateCreated, data.OofShard)
		if err != nil {
			fmt.Println("Ошибка отправки в БД:", err)
		}

		fmt.Println(result.RowsAffected())
		fmt.Println(result1.RowsAffected())
		fmt.Println(result2.RowsAffected()) // количество добавленных строк
		//fmt.Println("\n")
	}
}

func Find(db *sqlx.DB) ([]jsonstr.ModelJSON, error) {
	var results []jsonstr.ModelJSON
	//var js string
	rows, err := db.Queryx(`select * from Wildberry, Paymentt, Deliveryt`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		//var result jsonstr.ModelJSON
		result := jsonstr.ModelJSON{}
		var txt, a, b string
		//err = rows.StructScan(&result)
		err = rows.Scan(&result.OrderUID, &result.TrackNumber, &result.Entry, &a, &b, &txt, &result.Locale, &result.InternalSignature,
			&result.CustomerID, &result.DeliveryService, &result.Shardkey, &result.SmID, &result.DateCreated, &result.OofShard,
			&result.Payment.Transaction, &result.Payment.RequestID, &result.Payment.Currency, &result.Payment.Provider, &result.Payment.Amount,
			&result.Payment.PaymentDt, &result.Payment.Bank, &result.Payment.DeliveryCost, &result.Payment.GoodsTotal, &result.Payment.CustomFee,
			&result.Delivery.Name, &result.Delivery.Phone, &result.Delivery.Zip, &result.Delivery.City, &result.Delivery.Address, &result.Delivery.Region,
			&result.Delivery.Email) //не хотел так делать, пытался через gorm и sqlx, но не получилось исправить некоторые ошибки
		if err != nil {
			fmt.Println(err)
		}

		if err := json.Unmarshal([]byte(txt), &result.Items); err != nil {
			fmt.Println(err)

		}
		results = append(results, result)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}
	return results, err
}
