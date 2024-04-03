package main // Тема 2/6: Интерфейсы → Урок 2/5 "Интерфейсы в стандартной библиотеке"
// https://practicum.yandex.ru/trainer/go-basics/lesson/4f011178-b57f-4b95-8b47-5373fab03114/

import (
	"io"
	"log"
	"os"
	"strings"
)

// Подсказка: подумайте, как можно ограничить чтение.
// Нужно как-то подсчитывать и запоминать количество байт, оставшихся для чтения из reader.

type LimitedReader struct {
	reader io.Reader
	left   int //  запоминаем количество считанных байт
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{reader: r, left: n}
}

func (r *LimitedReader) Read(p []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if r.left < len(p) {
		p = p[0:r.left]
	}

	n, err := r.reader.Read(p)
	r.left -= n

	return n, err
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
