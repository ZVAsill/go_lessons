package main

import "fmt"

/*
1 почитать количество одинаковых цифр в слайсе

```go
func CountInts(nums []int) map[int]int {
	// implement me
}
```

на вход имеем  [ 2,3,4,5,3,2,2,4,5,6,7,5,3,2,2,4,5]

на выходе имеем мап где ключ уникальная цифра из слайса, значение количество ее повторений в слайсе

{

2: 5, // количество двоек 5 и т.д.

3: 3,

4: 3,

5: 4,

6: 1,

7: 1,

}
*/

func main() {
	var nums []int
	nums = []int{2, 3, 4, 5, 3, 2, 2, 4, 5, 6, 7, 5, 3, 2, 2, 4, 5}
	var count map[int]int
	count = CountInts(nums)
	fmt.Println(count)
}

func CountInts(nums []int) map[int]int {
	var c map[int]int
	c = make(map[int]int)
	for _, j := range nums {
		c[j]++
	}
	return c
}
