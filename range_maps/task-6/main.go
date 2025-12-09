package main

import "fmt"

/*
6 Сгруппировать всех юзеров по городу

	type User struct {
		Name string
		City: string
	}

	func GroupByCity(users []User) map[string][]User {
		// implement me
	}
*/
type User struct {
	Name string
	City string
}

func main() {
	var users []User
	users = []User{
		{"Василий", "Алматы"},
		{"Карлос", "Лас-Вегас"},
		{"Иван", "Москва"},
		{"Дмитрий", "Алматы"},
		{"Рикардо", "Марсель"},
		{"Илья", "Москва"},
		{"Дональд", "Вашингтон"},
		{"Ричард", "Лас-Вегас"},
	}
	var c map[string][]User
	c = GroupByCity(users)
	fmt.Println(c)
}

func GroupByCity(users []User) map[string][]User {
	var c map[string][]User
	c = make(map[string][]User)

	for i, j := range users {
		fmt.Println(i, j.City, j.Name)
		c[j.City] = append(c[j.City], j)
	}
	return c
}
