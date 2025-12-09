package main

import "fmt"

/*
4 Сделать инверсию мапы

```go

	Invert(m map[string]int) map[int]string {
		// implement me
	}

```

# Например имеем на входе

{

“a”: 1,

“b”: 2,

}

на выходе

{

1: “a”,

2: “b”,

}
*/
func main() {
	var m map[string]int
	var im map[int]string
	m = map[string]int{"a": 1, "b": 2}
	im = Invert(m)
	fmt.Println(im)
}

func Invert(m map[string]int) map[int]string {
	var im map[int]string
	im = make(map[int]string)
	for i, j := range m {
		im[j] = i
	}
	return im
}
