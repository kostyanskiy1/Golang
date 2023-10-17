package main

import (
	"fmt"
	"strings"
)

func uniq(str string) bool {
	str = strings.ToLower(str) // Приводим к нижнему регистру
	m := make(map[rune]bool)   // Создаем мапу для отслеживания уникальных символов
	for _, ch := range str {
		if m[ch] {
			return false // Если уже был, возвращаем false
		}
		m[ch] = true //иначе ставим true
	}
	return true
}

func main() {
	fmt.Println(uniq("abcd"))
	fmt.Println(uniq("abCdefAaf"))
	fmt.Println(uniq("aabcd"))
}
