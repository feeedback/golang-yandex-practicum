package main // Тема 2/4: Система типов → Урок 3/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/5d508ff8-07cf-4390-90d2-f9a502c5b133/

import "fmt"

const (
	one = 2*iota + 1 // укажите здесь формулу с iota
	three
	five
	seven
	nine
	eleven
)

func main() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
