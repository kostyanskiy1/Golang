package main

import "fmt"

func main() {
	old := "главрыба"
	new := ""

	for _, a := range old {
		b := string(a) + new
		b, new = new, b
	}
	fmt.Println(new)
}
