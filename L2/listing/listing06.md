Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main
import (
	"fmt"
)
func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}
func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Слайс сылается на область памяти, где лежит массив, поэтому при выполнении станет i[0] = "3"
Изменится и значение слайса в главной функции, но затем, при выполнении операции append(),
происходит создание нового слайса, который не будет возвращен из функции, и все дальнейшие операции 
после i[0] = "3" никак не повлияют на наш основной слайс. 
```