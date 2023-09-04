package main // Тема 1/4: Композитные типы → Урок 2/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/57d4b5d3-6df9-4653-a875-e83f47538e7a/

import (
	"fmt"
	"slices"
)

func reverse[T any](slice []T) []T {
	// // before Go 1.21.0
	// n := len(slice)
	// for i := 0; i < n/2; i++ {
	// 	slice[i], slice[n-1-i] = slice[n-1-i], slice[i]
	// }
	slices.Reverse(slice)

	return slice
}

func main() {
	input := "Привет, мир!"

	runes := []rune(input)

	runes = reverse(runes)

	fmt.Println(string(runes))
}
