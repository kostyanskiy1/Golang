package main

import "fmt"

func main() {

	mas := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mapa := make(map[int][]float64)
	for _, m := range mas {

		val := int(m/10) * 10 //шагом в 10 градусов
		fmt.Println(val)
		mapa[val] = append(mapa[val], m) //вставляем в срез по ключу
	}
	fmt.Println(mapa)
}
