package main

import (
	"bytes"
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func FindAnogram(arr *[]string) *map[string]*[]string {

	dict := make(map[string]struct{})
	for _, val := range *arr {
		val = string(bytes.ToLower([]byte(val)))
		dict[val] = struct{}{}
	}
	var set = make(map[string]*[]string)
	var buffSet = make(map[string][]string)

	for val := range dict {
		ss := SortString(val)
		buffSet[ss] = append(buffSet[ss], val)
	}

	fmt.Println("buffSet=", buffSet)
	for k, v := range buffSet {
		fmt.Println("k=", k, " v=", v)
		sort.Strings(buffSet[k])
		a := buffSet[k]
		fmt.Println("a=", a)
		set[v[0]] = &a
	}
	return &set
}

func main() {
	var strArr = []string{"тяпка", "пятак", "лиСток", "пЯтка", "слиТок", "столик", "стоЛик"}
	set := FindAnogram(&strArr)
	for k, v := range *set {
		fmt.Printf("%s: %v", k, v)
	}
}
