package main //ищем значение(в данном случае '0') в csv файле

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	const root = "D:/GOLANG/FILES/studing/tc.txt"
	//const root = "https://github.com/semyon-dev/stepik-go/blob/master/work_with_files/task_sep_1/task.data"
	file, err := os.Open(root)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	rs := bufio.NewReader(file)
	reader := csv.NewReader(rs)
	reader.Comma = ';'

	record, err := reader.Read()

	if err != nil && err != io.EOF {
		// io.EOF не совсем ошибка - это состояние, указывающее, что файл прочитан до конца
		fmt.Print(err)
	}

	// record - теперь это просто массив строк (array []string).  Проходим его посредством for и range:
	for num, item := range record {
		if item == "0" { //ищем 0 в csv файле
			fmt.Println("hell=", num+1)
			break
		}
	}

}
