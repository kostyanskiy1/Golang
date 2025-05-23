# WB Tech: level # 2 (Golang)
## Tasks
### Design patterns
Implement patterns, explain the applicability of each pattern, pros and cons, and real examples of using the pattern in practice.

Pattern "facade".
Pattern "builder".
Pattern "visitor".
Pattern "command".
Pattern "call chain".
Pattern "factory method".
Pattern "strategy".
Pattern "state".

### Development tasks
Programs must pass all tests. The code must pass go vet and golint checks.
Basic task

1. #### Basic task
Create a program that prints the exact time using the NTP library. Initialize as a go module. Use the github.com/beevik/ntp library. Write a program that prints the current time / the exact time using this library.

Requirements:
The program must be designed as a go module
The program must correctly handle library errors: output them to STDERR and return a non-zero exit code to the OS

2. #### Unpacking task
Create a Go function that performs primitive unpacking of a string containing repeating characters/runes, for example:
```"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (incorrect string)
"" => ""
```
Additional
Implement support for escape sequences.
For example:
```
qwe\4\5 => qwe45 (*)
qwe\45 ​​=> qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
```

If an invalid string was passed, the function should return an error. Write unit tests.

3. #### Utility sort
Sort lines in a file similar to the console utility sort (man sort - see the description and main parameters): the input is a file with unsorted lines, the output is a file with sorted ones.

Implement support for the following keys in the utility:

-k — specify a column for sorting (words in a line can act as columns, the default separator is a space)
-n — sort by numeric value
-r — sort in reverse order
-u — do not output duplicate lines

Additional

Implement support for the following keys in the utility:

-M — sort by month name
-b — ignore trailing spaces
-c — check if the data is sorted
-h — sort by numeric value taking into account suffixes

4. #### Search for anagrams in a dictionary

Write a function to search for all sets of anagrams in a dictionary.

For example:
'five-dollar coin', 'heel' and 'hoe' - belong to one set,
'leaflet', 'ingot' and 'table' - to another.

Requirements:
* Input data for the function: a reference to an array, each element of which is a word in Russian in utf8 encoding
* Output data: a reference to a map of anagram sets
* Key - the first word from the set encountered in the dictionary. Value - a reference to an array, each element of which is a word from the set.
* The array must be sorted in ascending order.
* Sets of one element must not be included in the result.
* All words must be converted to lower case.
* As a result, each word must occur only once.

5. #### grep utility

Implement a filtering utility similar to the console utility (man grep - see the description and main parameters).

Implement support for the following keys in the utility:
-A - "after" print +N lines after the match
-B - "before" print +N lines before the match
-C - "context" (A+B) print ±N lines around the match
-c - "count" (number of lines)
-i - "ignore-case" (ignore case)
-v - "invert" (instead of a match, exclude)
-F - "fixed", exact match with a line, not a pattern
-n - "line num", print the line number

6. #### The cut utility

Implement a utility similar to the console command cut (man cut). The utility should accept lines via STDIN, split them into columns by a separator (TAB), and output the requested ones.

Implement support for the following keys in the utility:
-f - "fields" - select fields (columns)
-d - "delimiter" - use a different delimiter
-s - "separated" - only lines with a delimiter

7. #### Or channel

Implement a function that will merge one or more done channels into a single channel if one of its constituent channels closes.

An obvious solution would be an expression using select that would implement this connection, but sometimes the total number of done channels that you work with at runtime is unknown. In this case, it is more convenient to use a call to a single function that, having accepted one or more or channels as input, would implement all the functionality.

Function definition:
```go
var or func(channels ...<- chan interface{}) <- chan interface{}
Example of using the function:
sig := func(after time.Duration) <- chan interface{} {
c := make(chan interface{})
go func() {
defer close(c)
time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
sig(2*time.Hour),
sig(5*time.Minute),
sig(1*time.Second),
sig(1*time.Hour),
sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
```

8. #### Interaction with OS

It is necessary to implement your own UNIX shell utility with support for a number of simple commands:

- cd <args> - change directory (this and that can be used as an argument)
- pwd - show the path to the current directory
- echo <args> - output argument to STDOUT
- kill <args> - "kill" the process passed as an argument (example: such-and-such example)
- ps - outputs general information on running processes in the format *such-and-such format*

It is also necessary to support the functionality of fork/exec commands

Additionally, it is necessary to support a pipeline on pipes (linux pipes, example cmd1 | cmd2 | .... | cmdN).

*A shell is a regular console program that, when launched, in an interactive session outputs a certain prompt
to STDOUT and waits for user input through STDIN. After waiting for input, it processes the command according to its logic
and, if necessary, displays the result on the screen. The interactive session is maintained until the exit command is entered (for example, \quit).

9. #### wget utility

Implement the wget utility with the ability to download entire sites.

10. #### telnet utility

Implement the simplest telnet client.

Examples of calls:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Requirements:
The program must connect to the specified host (ip or domain name + port) via TCP. After connecting, the program's STDIN should be written to the socket, and the data received from the socket should be output to STDOUT
Optionally, a timeout for connecting to the server can be passed to the program (via the --timeout argument, 10s by default)
When Ctrl+D is pressed, the program should close the socket and exit. If the socket is closed from the server side, the program should also exit. When connecting to a non-existent server, the program should exit via timeout

11. #### HTTP server

Implement an HTTP server for working with the calendar. As part of the task, it is necessary to work strictly with the standard HTTP library.

As part of the task, it is necessary to:
Implement helper functions for serializing domain objects to JSON.
Implement helper functions for parsing and validating the parameters of the /create_event and /update_event methods.
Implement HTTP handlers for each of the API methods using helper functions and domain objects.
Implement middleware for logging requests

API methods:

* POST /create_event
* POST /update_event
* POST /delete_event
* GET /events_for_day
* GET /events_for_week
* GET /events_for_month

Parameters are passed as www-url-form-encoded (i.e. regular user_id=3&date=2019-09-09). In GET methods, parameters are passed via queryString, in POST via the request body.
Each request should return a JSON document containing either {"result": "..."} in case of successful method execution, or {"error": "..."} in case of a business logic error.

As part of the task, it is necessary to:
Implement all methods.
Business logic should NOT depend on the HTTP server code.
In case of a business logic error, the server should return HTTP 503. In case of an input data error (invalid int, for example), the server should return HTTP 400. In case of other errors, the server should return HTTP 500. The web server should run on the port specified in the config and output each processed request to the log.





# WB Tech: level # 2 (Golang)
## Задания
### Паттерны проектирования
Реализовать паттерны, объяснить применимость каждого паттерна, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Паттерн «фасад». 
Паттерн «строитель». 
Паттерн «посетитель». 
Паттерн «комманда». 
Паттерн «цепочка вызовов». 
Паттерн «фабричный метод». 
Паттерн «стратегия». 
Паттерн «состояние». 


### Задачи на разработку
Программы должны проходить все тесты. Код должен проходить проверки go vet и golint.
Базовая задача

1. #### Базовая задача
Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

2. #### Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
```"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
```
Дополнительно 
Реализовать поддержку escape-последовательностей. 
Например:
```
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
```

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

3. #### Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

4. #### Поиск анаграмм по словарю

Написать функцию поиска всех множеств анаграмм по словарю. 


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
* Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
* Выходные данные: ссылка на мапу множеств анаграмм
* Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого, 
слово из множества.
* Массив должен быть отсортирован по возрастанию.
* Множества из одного элемента не должны попасть в результат.
* Все слова должны быть приведены к нижнему регистру. 
* В результате каждое слово должно встречаться только один раз.

5. #### Утилита grep


Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).


Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки

6. #### Утилита cut

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

7. #### Or channel 

Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.

Определение функции: 
```go
var or func(channels ...<- chan interface{}) <- chan interface{}
Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
```

8. #### Взаимодействие с ОС


Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение 
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике 
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).


9. #### Утилита wget

Реализовать утилиту wget с возможностью скачивать сайты целиком.

10. #### Утилита telnet

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123


Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

11. #### HTTP-сервер

Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.


В рамках задания необходимо:
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
Реализовать middleware для логирования запросов


Методы API: 
* POST /create_event 
* POST /update_event 
* POST /delete_event 
* GET /events_for_day 
* GET /events_for_week 
* GET /events_for_month


Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
Реализовать все методы.
Бизнес логика НЕ должна зависеть от кода HTTP сервера.
В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
