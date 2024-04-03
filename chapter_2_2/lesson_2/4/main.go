package main // Тема 2/4: Оператор отложенного вызова → Урок 2/3
// https://practicum.yandex.ru/trainer/go-basics/lesson/501ddace-d38d-41bc-8a56-02df538da473/

import (
	"fmt"
)

var Global = 5

func UseGlobal() {
	defer func(value int) {
		Global = value
	}(Global)
	// раннее связывание

	Global = 22

	fmt.Println(Global)
}

func main() {
	fmt.Println(Global)

	UseGlobal()

	fmt.Println(Global)
}
