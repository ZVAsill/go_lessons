package main

/*
### Задача 4. Полиморфизм: несколько реализаций одного интерфейса

**Условие:**

Добавь тип `Robot` с полем `ID string`, который тоже реализует `Greeter`, но приветствует так: `"Beep. I am robot <ID>"`.

В `main` передай в `SayHello` и `Person`, и `Robot`.
*/
import "fmt"

type Greeter interface {
	Greet() string
}

func (p Person) Greet() string {
	var s string
	s = "Hello. I am a human, my name is " + p.Name
	return s
}

type Person struct {
	Name string
}

func (r Robot) Greet() string {
	var s string
	s = "Beep. I am robot " + r.ID
	return s
}

type Robot struct {
	ID string
}

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	p := Person{Name: "Фёдор"}
	r := Robot{ID: "RX-78"}

	SayHello(p)
	SayHello(r)
}
