package main

import "fmt"

/*
Задание 5. Удаление дубликатов из слайса
Описание:Напиши функцию
func unique(nums []int) []int
которая возвращает новый слайс без дубликатов, сохраняя порядок элементов.
Пример:
Вход: [1 2 2 3 1 4 3]Выход: [1 2 3 4]
Подсказка:Используй map[int]bool для проверки уже встречавшихся элементов, и append() для формирования нового результата.
*/

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 3}
	var uniqueNums []int
	uniqueNums = unique(nums)
	fmt.Println(uniqueNums)
}

func unique(nums []int) []int {
	var uNums []int
	var isHasuNum map[int]bool
	isHasuNum = make(map[int]bool)
	for _, j := range nums {
		if isHasuNum[j] != true {
			uNums = append(uNums, j)
			isHasuNum[j] = true
		}
	}

	return uNums
}
