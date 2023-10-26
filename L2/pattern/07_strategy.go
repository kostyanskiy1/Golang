package pattern

import (
	"fmt"
)

// реализация интерфейса для сортирвки.
type StrategySort interface {
	Sort([]int)
}

// bubble sort.
type BubbleSort struct {
}

// Sort sorts data.
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// insertion sort.
type InsertionSort struct {
}

// Sort sorts data.
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// Контекст обеспечивает условия для реализации стратегии.
type Context struct {
	strategy StrategySort
}

// алгоритм замены стратегии.
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// сортирует данные выбранной стратегией.
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}

func StrategyFunc() {

	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}

	ctx := new(Context)

	ctx.Algorithm(&BubbleSort{})

	ctx.Sort(data1)

	ctx.Algorithm(&InsertionSort{})

	ctx.Sort(data2)

	fmt.Println("data1:", data1)
	fmt.Println("data2:", data2)
}
