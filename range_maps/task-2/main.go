package main

import "fmt"

/*
2 Посчитать количество слов в строке

```go

	WordCount(s string) map[string]int {
		 // implement me
	}

```

# Аналогично CountInts, только ключ строка

на вход имеем “Эх, если бы да кабы, во рту грибы росли, то был бы не рот, а целый огород”

на выходе

{

“Эх”: 1,

“если”: 1,

“бы”: 2, // совпадение 2 раза

“да”: 1,

“кабы”: 1,

“во”: 1,

“рту”: 1,

“грибы”: 1,

“росли”: 1,

“то”: 1,

“был”: 1,

“не”: 1,

“рот”: 1,

“а”: 1,

“целый”: 1,

“огород”: 1,

}
*/
func main() {
	var s string
	s = "Эх, если бы да кабы, во рту грибы росли, то был бы не рот, а целый огород"
	var count map[string]int
	count = WordCount(s)
	fmt.Println(count)
}

func WordCount(s string) map[string]int {
	var c map[string]int
	c = make(map[string]int)
	word := ""
	for _, j := range s {
		if isLetter(j) {
			word = word + string(j)
		} else {
			if word != "" {
				c[word]++
				word = ""
			}
		}
	}
	return c
}

func isLetter(r rune) bool {
	return (r >= 'А' && r <= 'я') || r == 'Ё' || r == 'ё'
}
