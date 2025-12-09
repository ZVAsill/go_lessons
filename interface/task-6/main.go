package main

/*
### Задача 6. Указатель/значение и интерфейс

**Условие:**

Создай интерфейс `Updater` с методом `Update(msg string)`.

Создай тип `Logger` с полем `Last string`.

Сделай метод `Update` с **указательным** получателем, который сохраняет `msg` в `Last`.

Попробуй присвоить в переменную типа `Updater` и `Logger`, и `*Logger`. Посмотри, что компилируется.
*/
import "fmt"

type Updater interface {
	Update(msg string)
}

func (l *Logger) Update(msg string) {
	l.Last = msg
}

type Logger struct {
	Last string
}

func main() {
	//var u1 Updater
	var u2 Updater

	l := Logger{}

	//u1 = l  // <- так не компилируется
	u2 = &l // <- а так компилируется

	u2.Update("hello")
	fmt.Println("Last:", l.Last)
}
