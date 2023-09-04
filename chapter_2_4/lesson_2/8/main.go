package main // Тема 1/4: Композитные типы → Урок 2/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/57d4b5d3-6df9-4653-a875-e83f47538e7a/

import "fmt"

func main() {
	input := "The quick brown 狐 jumped over the lazy 犬"

	runes := []rune(input)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	str := string(runes)
	fmt.Println(str)
}
