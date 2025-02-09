# go-core-task

### Задания по модулю Go-core

Создать в гите директорию go-core-task. Каждое задание делать в отдельной верке этой директории и потом слить все задания в master.
Структура дирректории должна выглядеть примерно так:
```
go-core-task/
├── 1/
│     ├── main_1.go
│     └── main1_test.go
├── 2/
│     ├── main_2.go
│     └── main2_test.go
...
├── 9/
│     ├── main_9.go
│     └── main9_test.go
└── README.md
```

Выполнив задания, ссылку на репозиторий присылать своему руководителю направления (Golang)

---

### Задание 1
Напишите программу на Go, которая:

1. Создает несколько переменных различных типов данных:
```   
int (три числа в десятичной, восьмеричной и шеснадцатиричной системах)
float64
string
bool
complex64
```
2. Определяет тип каждой переменной и выводит его на экран.
3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
4. Преобразовать эту строку в срез рун.
5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.

* Напишите unit тесты к созданным функциям

Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

Входные числа из пункта 1 могут быть:
```
var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатиричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"         // Тип string
var isActive bool = true           // Тип bool
var complexNum complex64 = 1 + 2i  // Тип complex64
```

---

#### Задание 2

1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
образом (при каждом запуске должны получаться новые значения)

2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.

3. Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.

4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.

5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.

6. Напишите main функцию, в которой протестируете все вышеописанные функции. Выведите результаты на экран.

* Напишите unit тесты к созданным функциям

Примечание.
В качестве originalSlice можно использовать ```originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}```

---

#### Задание 3

Реализуйте структуру данных StringIntMap, которая будет использоваться для хранения пар "строка - целое число". Ваша
структура должна поддерживать следующие операции:

1. Добавление элемента: Метод Add(key string, value int), который добавляет новую пару "ключ-значение" в карту.

2. Удаление элемента: Метод Remove(key string), который удаляет элемент по ключу из карты.

3. Копирование карты: Метод Copy() map[string]int, который возвращает новую карту, содержащую все элементы текущей карты.

4. Проверка наличия ключа: Метод Exists(key string) bool, который проверяет, существует ли ключ в карте.

5. Получение значения: Метод Get(key string) (int, bool), который возвращает значение по ключу и булевый флаг, указывающий на успешность операции.

* Напишите unit тесты к созданным функциям

---

#### Задание 4

На вход подаются два неупорядоченных слайса строк.
Например:
```
slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
slice2 := []string{"banana", "date", "fig"}
```
Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.

* Напишите unit тесты к созданным функциям

---

#### Задание 5

На вход подаются два неупорядоченных слайса int любой длины.
Например:
```
a := []int{65, 3, 58, 678, 64}
b := []int{64, 2, 3, 43}
```
Напишите функцию, которая проверяет, есть ли пересечения значений между двумя слайсами и возвращает:
- bool значение есть ли хотя бы одно пересечение в значениях входных срезов
- срез []int с пересеченными значениями (если таких значений нет, то возвращать пустой срез).
Т. е. если взать слайсы a и b из премере, то вернет ```true, []int{64, 3}```.

* Напишите unit тесты к созданным функциям

---

#### Задание 6

Напишите генератор случайных чисел используя небуфферизированный канал.

* Напишите unit тесты к созданным функциям

---

#### Задание 7

Напишите программу на Go, которая сливает N каналов в один.

* Напишите unit тесты к созданным функциям

---

#### Задание 8

Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

* Напишите unit тесты к созданным функциям

---

#### Задание 9

Сделать конвейер чисел
Даны два канала.
В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления,
затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.

Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

* Напишите unit тесты к созданным функциям

---
