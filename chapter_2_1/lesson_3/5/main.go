package main // Тема 1/4: Композитные типы → Урок 3/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/9f07b681-ffbe-4835-91ad-48341ca291da/

import "fmt"

func RemoveDuplicates(input []string) []string {
	mapWordToExist := make(map[string]bool)
	uniqWords := []string{}

	for _, word := range input {
		if !mapWordToExist[word] {
			mapWordToExist[word] = true

			uniqWords = append(uniqWords, word)
		}
	}

	return uniqWords
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	unique := RemoveDuplicates(input)
	fmt.Println(unique)
}
