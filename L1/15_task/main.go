package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] //ссылаемся на строку которую, не используем, и сборщик мусора ее убрать не может

	justString = string(v[:100]) //1 Связь с базовым массивом теряется
	copy(justString, v[:100])    //2 Связь с базовым массивом теряется
}

func main() {
	someFunc()
}
