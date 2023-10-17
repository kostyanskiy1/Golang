package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1, 6, 5, 3, 4}
	sort.Ints(nums) // Используем функцию sort.Ints() для сортировки массива
	fmt.Println(nums)
}
