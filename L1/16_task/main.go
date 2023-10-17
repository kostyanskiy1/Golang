package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1, 6, 5, 3, 4}

	sort.Slice(nums, func(i, j int) bool { //Сортировка слайсов
		return nums[i] < nums[j]
	})

	fmt.Println(nums) // [1 2 3 4 5 6]
}
