package main // Тема 1/4: Композитные типы → Урок 2/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/57d4b5d3-6df9-4653-a875-e83f47538e7a/

import (
	"fmt"
	"slices"
)

func reverse[T any](slice []T) []T {
	// // before Go 1.21.0
	// n := len(slice)
	// for i := 0; i < n/2; i++ {
	// 	slice[i], slice[n-1-i] = slice[n-1-i], slice[i]
	// }
	slices.Reverse(slice)

	return slice
}

func main() {
	const max = 100
	const head = 10
	const tail = 10

	nums := make([]int, 0, max)

	for i := 0; i < max; i++ {
		nums = append(nums, i+1)
	}

	nums = append(nums[:head], nums[max-tail:]...)

	nums = reverse(nums)

	fmt.Println(nums)
}
