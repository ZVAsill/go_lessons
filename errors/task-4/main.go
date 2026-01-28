package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
## 4) `errors.As` + типизированная ошибка (полезные поля)

**Цель:** делать ошибки, которые можно “распаковать”.

**Задание:** создай тип:

```go

	type ValidationErrorstruct {
		Fieldstring
		Reasonstring
	}

```

- Реализуй `Error() string`
- Реализуй `ValidateSignup(email, password string) error`
  - email должен содержать `@`
  - password должен быть минимум 8 символов
  - на первое найденное нарушение возвращай `ValidationError`

**Готово, если:**

- в коде можно сделать:
  - `var ve *ValidationError; errors.As(err, &ve)` → `true`
  - `ve.Field` правильно заполнен
*/
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid field : %s ; reason: %s", e.Field, e.Reason)
}

func main() {
	var ve *ValidationError
	email := "dgdg@fd.ru"
	password := "123456"
	err := ValidateSignup(email, password)
	if err != nil {
		errors.As(err, &ve)
		fmt.Println("ошибка полностью: " + err.Error())
		fmt.Println("распакованная ошибка название поля: " + ve.Field)
		fmt.Println("распакованная ошибка причина: " + ve.Reason)
	} else {
		fmt.Println("Validation success")
	}
}

func ValidateSignup(email, password string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{"email", "email is invalid"}
	}
	if utf8.RuneCountInString(password) < 8 {
		return &ValidationError{"password", "password length less 8 symbols"}
	}
	return nil
}
