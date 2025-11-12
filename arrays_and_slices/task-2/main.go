package main

/*
Задание 2. Умножение всех элементов слайса
Описание:Создай слайс из 5 чисел. Реализуй функцию multiply(s []int, factor int) []int, которая умножает каждый элемент на заданный множитель.
Пример:
Вход: [2 4 6], factor = 3  Выход: [6 12 18]
Подсказка:Слайс изменяется по ссылке, если его не копировать. Попробуй обе версии: с изменением на месте и с копией (copy()).
*/
func main() {
	s := []int{1, 2, 3, 4, 5}
	multiply(s, 3)
	println("----multiply --------")
	for i := 0; i < len(s); i++ {
		println(s[i])
	}
	println("-----multiply copy-------")
	s2 := multiply_copy(s, 3)
	for i := 0; i < len(s2); i++ {
		println(s2[i])
	}
	println("------------")

}

func multiply(s []int, factor int) {
	for i := 0; i < len(s); i++ {
		s[i] *= factor
	}
}

func multiply_copy(s []int, factor int) []int {
	copy_slice := make([]int, len(s))
	copy(copy_slice, s)
	for i := 0; i < len(copy_slice); i++ {
		copy_slice[i] *= factor
	}
	return copy_slice
}
