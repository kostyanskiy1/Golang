# WB Tech: level # 1 (Golang)
## How to complete tasks
The tasks do not include verbal solutions — only code. One solution — one file with well-commented code. Each solution or impossibility of a solution must be explained.

The use of any reference resources, the involvement of third-party experts, etc., etc. is allowed and encouraged.

The main evaluation criterion is a clear understanding of “how it works”. Some tasks can be solved in several ways, in which case it is necessary to provide the maximum possible number of options.

You can ask questions both about the conditions of the tasks and about their solution. The ideal option is to demonstrate your solutions and get maximum feedback from experienced Wildberries developers.
## Tasks
1. A Human structure is given (with an arbitrary set of fields and methods). Implement the embedding of methods in the Action structure from the parent Human structure (analogous to inheritance).

2. Write a program that will competitively calculate the value of the squares of numbers taken from an array (2,4,6,8,10) and output their squares to stdout.

3. Given a sequence of numbers: 2,4,6,8,10. Find the sum of their squares (22+32+42….) using competitive calculations.

4. Implement constant data writing to a channel (main thread). Implement a set of N workers that read arbitrary data from the channel and output to stdout. It is necessary to be able to select the number of workers at startup. The program must terminate by pressing Ctrl+C. Choose and justify a method for terminating the work of all workers.

5. Develop a program that will sequentially send values ​​to a channel, and read from the other side of the channel. After N seconds, the program must terminate.

6. Implement all possible ways to stop the execution of a goroutine.

7. Implement concurrent data writing to map.

8. Given an int64 variable. Develop a program that sets the i-th bit to 1 or 0.

9. Develop a number pipeline. Given two channels: the first one writes numbers (x) from the array, the second one writes the result of the x*2 operation, after which the data from the second channel should be output to stdout.

10. Given a sequence of temperature fluctuations: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Combine these values ​​into groups with a step of 10 degrees. The sequence in the subsets is not important.

Example: ```-20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, ​​20: {24.5}, etc.```

11. Implement the intersection of two unordered sets.

12. There is a sequence of strings - (cat, cat, dog, cat, tree) create its own set for it.

13. Swap two numbers without creating a temporary variable.

14. Develop a program that can determine the type of a variable at runtime: int, string, bool, channel from a variable of type interface{}.

15. What negative consequences can this code fragment lead to, and how can this be fixed? Provide a correct example of implementation.

```go
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}

func main() {
someFunc()
}
```

16. Implement quicksort of an array using built-in language methods.

17. Implement binary search using built-in language methods.

18. Implement a counter structure that will be incremented in a concurrent environment. Upon completion, the program should output the final value of the counter.

19. Develop a program that reverses the string given to it (for example: "glavryba — abyrvalg"). The symbols can be unicode.

20. Develop a program that reverses words in a string.

Example: "snow dog sun — sun dog snow".

21. Implement the "adapter" pattern on any example.

22. Develop a program that multiplies, divides, adds, subtracts two numeric variables a, b, the value of which is > 2^20.

23. Remove the i-th element from the slice.

24. Develop a program to find the distance between two points, which are represented as a Point structure with encapsulated parameters x, y and a constructor.

25. Implement your own sleep function.

26. Develop a program that checks that all characters in a string are unique (true if unique, false etc.). The check function must be case-insensitive.

For example:
```go
abcd — true
abCdefAaf — false
aabcd — false
```



# WB Tech: level # 1 (Golang)
## Как делать задания
В заданиях никаких устных решений — только код. Одно решение — один файл с хорошо откомментированным кодом. Каждое решение или невозможность решения надо объяснить.

Разрешается и приветствуется использование любых справочных ресурсов, привлечение сторонних экспертов и т.д. и т.п. 


Основной критерий оценки — четкое понимание «как это работает». Некоторые задачи можно решить несколькими способами, в этом случае требуется привести максимально возможное количество вариантов.

Можно задавать вопросы, как по условию задач, так и об их решении. Идеальный вариант — продемонстрировать свои решения и получить максимальный фидбэк от опытных разработчиков Wildberries.
## Задания
1. Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

2. Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.


3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.


4. Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте. Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.



5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.


6. Реализовать все возможные способы остановки выполнения горутины. 


7. Реализовать конкурентную запись данных в map.


8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.


9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.


10. Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.


Пример: ```-20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.```


11. Реализовать пересечение двух неупорядоченных множеств.


12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.


13. Поменять местами два числа без создания временной переменной.


14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.


15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

```go
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
```

16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.


17. Реализовать бинарный поиск встроенными методами языка.


18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.


19. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.


20. Разработать программу, которая переворачивает слова в строке. 
Пример: «snow dog sun — sun dog snow».


21. Реализовать паттерн «адаптер» на любом примере.


22. Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.


23. Удалить i-ый элемент из слайса.


24. Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.


25. Реализовать собственную функцию sleep.


26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например: 
```go
abcd — true
abCdefAaf — false
aabcd — false
```
