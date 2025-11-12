package main

/*
Задание 3. Обрезка и копирование
Описание:Создай слайс numbers := make([]int, 0, 20) и добавь в него числа 1–10.Затем создай под-слайс part := numbers[3:8].Создай копию copyPart, чтобы освободиться от исходного массива.Измени элементы в part и убедись, что это не влияет на copyPart.
Выведи:
длину и ёмкость part и copyPart
адрес первого элемента (через &part[0]) для сравнения

Цель:Понять, когда слайсы указывают на одну память, а когда нет.
*/

func main() {
	numbers := make([]int, 0, 20)
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i+1)
	}
	println("--numbers start--")
	for i := 0; i < 10; i++ {
		println(numbers[i])
	}
	println("--numbers end--")
	part := numbers[3:8]
	println("--part start--")
	for i := 0; i < len(part); i++ {
		println(part[i])
	}
	println("--part end--")
	copyPart := make([]int, len(part))
	copy(copyPart, part)
	println("--copyPart start--")
	for i := 0; i < len(copyPart); i++ {
		println(copyPart[i])
	}
	println("--copyPart end--")

	part[0] = 12
	println("--part edited start--")
	for i := 0; i < len(part); i++ {
		println(part[i])
	}
	println("--part edited end--")

	println("--copyPart again start--")
	for i := 0; i < len(copyPart); i++ {
		println(copyPart[i])
	}
	println("--copyPart again end--")

	println("part len", len(part))
	println("copyPart len", len(copyPart))
	println("part cap", cap(part))
	println("copyPart cap", cap(copyPart))

	println("part[0] address", &part[0])
	println("copyPart[0] address", &copyPart[0])
}
