package main

/*
Задание 1. Сумма элементов массива
Описание:Создай массив из 10 целых чисел (от 1 до 10). Напиши функцию sum(arr [10]int) int, которая возвращает сумму всех элементов.
Пример:
Вход: [1 2 3 4 5 6 7 8 9 10]  Выход: 55
Подсказка:Используй цикл for i := 0; i < len(arr); i++ { ... }
*/
func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sum(a)
	println(sum)
}

func sum(array [10]int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}
