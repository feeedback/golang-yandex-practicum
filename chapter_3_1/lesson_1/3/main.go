package main // Тема 1/6: ООП → Урок 1/3 "Методы"
// https://practicum.yandex.ru/trainer/go-basics/lesson/59203fa8-851b-444a-82f5-d031138079d2/

import "fmt"

type MyInt int

func (my MyInt) Set(v int) {
	my = MyInt(v)
}

func main() {
	var myInt MyInt

	myInt.Set(5)

	fmt.Println(myInt) // 0
}
