package main

import "fmt"

/*
3 Проверить, является ли слово Анаграммой

```go
func IsAnagram(a, b string) bool {
	// implement me
}
```

например на вход имеем (”кабан”, “банка”) возвращаем true
*/

func main() {
	var a string
	var b string
	a = "кабан"
	b = "банка"
	var c bool
	c = IsAnagram(a, b)
	fmt.Println(c)
}

func IsAnagram(a, b string) bool {
	var c bool
	var aMap, bMap map[rune]int
	aMap = make(map[rune]int)
	bMap = make(map[rune]int)
	c = true

	for _, j := range a {
		aMap[j]++
	}

	for _, j := range b {
		bMap[j]++
	}

	if len(a) != len(b) {
		c = false
	}

	for i, va := range aMap {
		vb := bMap[i]
		if va != vb {
			c = false
		}
	}

	return c
}
