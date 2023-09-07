package main // Тема 1/4: Композитные типы → Урок 3/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/9f07b681-ffbe-4835-91ad-48341ca291da/

func main() {}

func find(nums []int, queryNum int) []int {
	mapValueToIdx := make(map[int]int)

	for idx, val := range nums {
		diff := queryNum - val

		if idx2, ok := mapValueToIdx[diff]; ok {
			return []int{idx, idx2}
		}

		mapValueToIdx[val] = idx
	}

	return nil
}
