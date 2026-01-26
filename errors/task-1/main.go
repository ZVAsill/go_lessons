package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
## 1) валидатор

**Цель:** научиться возвращать ошибки и не усложнять.

**Задание:** реализуй функции:

- `ParsePositiveInt(s string) (int, error)`
    - парсит `s` в int
    - если `s` не число → верни ошибку от `strconv.Atoi` (как есть)
    - если число `<= 0` → верни `errors.New("must be positive")`
- `SumTwo(a, b string) (int, error)`
    - вызывает `ParsePositiveInt` для `a` и `b`
    - если ошибка → сразу возвращает её
    - иначе возвращает сумму

**Готово, если:**

- `SumTwo("10","20") == 30`
- `SumTwo("x","20")` возвращает ошибку от `Atoi`
- `SumTwo("-1","20")` возвращает `"must be positive"`
*/

func main() {
	n, err := SumTwo("-1", "20")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func ParsePositiveInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if num < 0 {
		return 0, errors.New("must be positive")
	}
	return num, nil
}

func SumTwo(a, b string) (int, error) {
	a1, err := ParsePositiveInt(a)
	if err != nil {
		return 0, err
	}
	b1, err := ParsePositiveInt(b)
	if err != nil {
		return 0, err
	}
	return a1 + b1, nil
}
