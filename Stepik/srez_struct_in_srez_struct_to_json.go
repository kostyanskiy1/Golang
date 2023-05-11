package main

import (
	"encoding/json"
	"fmt"
)

//	Затем создал экземпляр:

type Stud struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int //срез структуры в срезе cтруктуры
}
type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Stud
}
type answer struct {
	Average float64
	Max     int
}

func main() {

	abc := Group{
		ID: 123,
		Students: []Stud{
			{"asda", "asdasd", "12w313", "as414dasd", "asdaca", "asdadqweasd", []int{1, 4, 57, 123, 35, 123}},
			{"", "", "", "", "", "", []int{2, 2, 2}},
			{"", "", "", "", "", "", []int{5, 3, 3}},
		},
	}

	kol := 0
	count := 0
	max := 0
	for _, res := range abc.Students {

		kol++
		for _, val := range res.Rating {
			count++
			if max <= val { //здесь считает мах

				max = val
			}

		}
	}
	diff := float64(count) / float64(kol) //здесь считает именно среднее количество оценок на каждого студента

	//fmt.Printf("%.1f", dif)
	ass := answer{
		Average: diff,
		Max:     max,
	}

	fmt.Println(ass)
	data, err := json.MarshalIndent(ass, "", "    ")
	if err != nil {
		fmt.Println(err)

	}

	fmt.Printf("%s", data)
}
