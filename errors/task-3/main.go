package main

import (
	"errors"
	"fmt"
)

/*
## `errors.Is` и “сентинел”: поиск пользователя

**Цель:** научиться отличать “не найдено” от “прочих”.

**Дано:**

```go
var ErrNotFound = errors.New("not found")

```

**Задание:** реализуй:

- `FindUserID(name string) (int, error)`
  - если `name == ""` → `fmt.Errorf("empty name: %w", ErrNotFound)` (да, тут специально)
  - если `name == "admin"` → `return 1, nil`
  - иначе → `return 0, fmt.Errorf("user %q: %w", name, ErrNotFound)`

- `IsNotFound(err error) bool` через `errors.Is(err, ErrNotFound)`

**Готово, если:**

- `IsNotFound(err)` true для пустого имени и неизвестного имени
- `admin` не даёт ошибку
*/
var ErrNotFound = errors.New("not found")

func main() {
	name := ""
	m, err := FindUserID(name)
	if err != nil {
		fmt.Printf("Строка запроса поиска пользователя:\"%s\"; Результат для IsNotFound(err) = ", name)
		if IsNotFound(err) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else {
		if m == 1 {
			fmt.Printf("Пользователь : \"%s\" найден\n", name)
		}
	}
}

func FindUserID(name string) (int, error) {
	if name == "" {
		return 0, fmt.Errorf("empty name: %w", ErrNotFound)
	} else if name == "admin" {
		return 1, nil
	} else {
		return 0, fmt.Errorf("user %q: %w", name, ErrNotFound)
	}
}

func IsNotFound(err error) bool {
	if errors.Is(err, ErrNotFound) {
		return true
	}
	return false
}
