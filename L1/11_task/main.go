package main

import "fmt"

func calc(mas1, mas2 []int) []int {
	res := make([]int, 0, 0)

	for _, m1 := range mas1 {
		for _, m2 := range mas2 {
			if m1 == m2 {
				res = append(res, m1)
			}
		}
	}

	return res
}

func main() {
	mas1 := []int{1, 2, 3, 4, 5}
	mas2 := []int{6, 7, 8, 3, 5, 5, 1}

	res := calc(mas1, mas2)
	fmt.Println("Пересечение:", res)
}
