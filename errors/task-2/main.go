package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
## 2) Контекст + обёртки

**Цель:** правильно добавлять контекст через `%w` и не терять причину.

**Задание:** реализуй:

- `ReadFileTrim(path string) (string, error)`
  - читает файл `os.ReadFile`
  - если ошибка → `fmt.Errorf("read %q: %w", path, err)`
  - возвращает строку `strings.TrimSpace(...)`

**Проверки:**

- при отсутствующем файле `errors.As(err, *os.PathError)` должно быть `true`
- сообщение должно содержать `read "..."`
*/
func main() {
	s, err := ReadFileTrim("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

func ReadFileTrim(path string) (string, error) {
	f, err := os.ReadFile(path)
	var pathErr *os.PathError
	if err != nil {
		if errors.As(err, &pathErr) {
			return "read...", nil
		} else {
			return "", fmt.Errorf("read %q: %w", path, err)
		}
	}
	return strings.TrimSpace(string(f)), nil
}
