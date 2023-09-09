package main // Тема 2/4: Функции → Урок 1/3
// https://practicum.yandex.ru/trainer/go-basics/lesson/97a43227-9d61-49b4-a4e8-10c1110aa2e4/

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	PrintAllFilesWithFilter(".", "lesson")
	PrintAllFilesWithFilter(".", "Functions")
}

func PrintAllFilesWithFilter(path string, filter string) {
	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, file := range files {
		filename := filepath.Join(path, file.Name())

		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}

		if file.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}
