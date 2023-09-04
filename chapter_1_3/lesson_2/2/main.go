package main // Тема 3/4: Управление потоком выполнения → Урок 2/3
// https://practicum.yandex.ru/trainer/go-basics/lesson/efe00720-0547-4219-98d1-68c0cd3f9780/

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		isDivisibleBy3 := i%3 == 0
		isDivisibleBy5 := i%5 == 0

		switch {
		case isDivisibleBy3 && isDivisibleBy5:
			fmt.Println("Fizz" + "Buzz")
		case isDivisibleBy3:
			fmt.Println("Fizz")
		case isDivisibleBy5:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
