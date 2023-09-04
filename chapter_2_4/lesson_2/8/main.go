package main // Тема 1/4: Композитные типы → Урок 2/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/57d4b5d3-6df9-4653-a875-e83f47538e7a/

import (
	"fmt"
	"slices"
)

func main() {
	input := "Привет, мир!"

	runes := []rune(input)

	// // before Go 1.21.0
	// n := len(runes)
	// for i := 0; i < n/2; i++ {
	// 	runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	// }
	slices.Reverse(runes)

	fmt.Println(string(runes))
}
