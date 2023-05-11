package main ///чтение файла и поиск

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

// Я создал следующие типы структур:
type dataStruct struct {
	GlobalId int64 `json:"global_id"` //поиск только в определенном столбце
}

type rawData struct {
	units []dataStruct //  срез структур:
}

func main() {
	var r rawData

	rs, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)

	//data, errReadAll := io.ReadAll(rs)
	if err := json.Unmarshal(body, &r.units); err != nil {
		fmt.Println(err)
		return
	}

	var count int64

	for _, res := range r.units {
		count += res.GlobalId //считываем значения

	}
	fmt.Println(count)
}
