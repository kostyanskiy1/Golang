package main

import "fmt"

func Set(mas []string) map[string]struct{} { //map[]struct{} экономит место
	set := make(map[string]struct{})
	for _, str := range mas {
		set[str] = struct{}{}
	}
	return set
}

func main() {
	mas := []string{"cat", "cat", "dog", "cat", "tree"}
	mySet := Set(mas) //собственное множество
	fmt.Println(mySet)
}
