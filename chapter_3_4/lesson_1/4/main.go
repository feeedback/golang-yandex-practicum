package main // Тема 4/6: Тестирование → Урок 1/3 "Тестирование"
// https://practicum.yandex.ru/trainer/go-basics/lesson/d01326ef-d2a3-46b1-8dd1-f4a7f6fa8d0e/

func EstimateValue(value int) string {
	switch {
	case value < 10:
		return "small"

	case value < 100:
		return "medium"

	default:
		return "big"
	}
}
