package main // Тема 1/4: Композитные типы → Урок 3/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/9f07b681-ffbe-4835-91ad-48341ca291da/

import (
	"fmt"
)

func main() {
	pricelist := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	const maxPrice = 500
	var delicates []string

	for key, val := range pricelist {
		if val > maxPrice {
			delicates = append(delicates, key)
		}
	}

	fmt.Println("Перечень деликатесов:")
	for _, val := range delicates {
		fmt.Println(val)
	}

	sum := 0
	for _, val := range order {
		sum += pricelist[val]
	}
	fmt.Println("Стоимость заказа ", sum)
}
