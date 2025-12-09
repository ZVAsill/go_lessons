package main

/*
### Задача 1. Простейший метод со значением-получателем

**Условие:**

Создай структуру `Point` с полями `X` и `Y` (int).

Добавь метод `DistanceFromOrigin() int`, который возвращает расстояние до начала координат по формуле `|X| + |Y|` (манхэттен).

В `main` создай точку и выведи расстояние.
*/
import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) DistanceFromOrigin() int {
	var zero int
	zero = AbsInt(p.X) + AbsInt(p.Y)
	return zero
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	p := Point{X: -3, Y: 5}
	fmt.Println("Distance:", p.DistanceFromOrigin())
}
