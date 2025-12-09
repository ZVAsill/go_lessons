package main

/*
### Задача 5. Срез интерфейсов

**Условие:**

Используя интерфейс `Greeter` и типы `Person` и `Robot`, сделай срез `[]Greeter` и в цикле вызови `Greet()` у каждого элемента.
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

func main() {
	var greeters []Greeter
	greeters = append(greeters, Person{"Василий"}, Robot{"R2D2"}, Person{"Хоттабыч"}, Robot{"Curiosity"}, Robot{"ВАЛЛ-И"})

	for _, g := range greeters {
		fmt.Println(g.Greet())
	}
}
