package main // Тема 2/4: Функции → Урок 1/3
// https://practicum.yandex.ru/trainer/go-basics/lesson/97a43227-9d61-49b4-a4e8-10c1110aa2e4/

import (
	"fmt"
	"os"
	fpath "path/filepath"
	"strings"
)

func main() {
	PrintFilesWithFuncFilter(".", func(filepath string) bool { return strings.Contains(filepath, "main") })

	PrintFilesWithFuncFilter(".", func(filepath string) bool { return strings.HasPrefix(fpath.Base(filepath), "c") })
}

func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	// создаём переменную, содержащую функцию обхода
	// мы создаём её заранее, а не через оператор :=, чтобы замыкание могло сослаться на него
	var walk func(string)

	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			filepath := fpath.Join(path, f.Name())

			if predicate(filepath) {
				fmt.Println(filepath)
			}

			if f.IsDir() {
				walk(filepath)
			}
		}
	}

	walk(path)
}
