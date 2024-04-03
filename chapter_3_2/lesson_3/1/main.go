package main // Тема 2/6: Интерфейсы → Урок 3/5 "Пустой интерфейс и приведение типов"
// https://practicum.yandex.ru/trainer/go-basics/lesson/c752e97d-d8da-48bc-8506-669475204e2e/

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch va := a.(type) {
	case int:
		return va * b

	case string:
		return strings.Repeat(va, b)

	case fmt.Stringer:
		return strings.Repeat(va.String(), b)

	default:
		return nil
	}
}
