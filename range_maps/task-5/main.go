package main

import "fmt"

/*
5 слияние 2х мап

```go
MergeMaps(dst, src map[string]int) map[string]int {
	// implement me
}
```

на вход имеем {”a”:3, “c”:5, ”f”: 3}, {”f”:2, “d”: 4, “a”: 5}

на выход получаем {”a”:8, “c”: 5, “f”: 5, “d”: 4}
*/

func main() {
	var m1, m2, m3 map[string]int
	m1 = map[string]int{"a": 3, "c": 5, "f": 3}
	m2 = map[string]int{"f": 2, "d": 4, "a": 5}
	m3 = MergeMaps(m1, m2)
	fmt.Println(m3)
}

func MergeMaps(dst, src map[string]int) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	for i, j := range dst {
		m[i] = m[i] + j
	}
	for i, j := range src {
		m[i] = m[i] + j
	}

	return m
}
