package main // Тема 3/6: Обработка ошибок → Урок 1/3 "Ошибки"
// https://practicum.yandex.ru/trainer/go-basics/lesson/93901ecb-3249-446c-8e19-ed9bd878c062/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type SliceError []error

// Статически создаём ошибки
var (
	ErrLineTooLong              = errors.New(`line is too long`)
	ErrLineIncludesNumbers      = errors.New(`found numbers`)
	ErrLineNotIncludesTwoSpaces = errors.New(`no two spaces`)
)

func (errs SliceError) Error() string {
	var errorStrings []string

	for _, err := range errs {
		errorStrings = append(errorStrings, err.Error())
	}
	return strings.Join(errorStrings, ";")
}

func MyCheck(input string) error {
	var (
		errList  SliceError
		spaces   int
		hasDigit bool
	)

	if len([]rune(input)) >= 20 {
		errList = SliceError{ErrLineTooLong}
	}

	for _, ch := range input {
		if ch == ' ' {
			spaces++
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}

	if hasDigit {
		errList = append(errList, ErrLineIncludesNumbers)
	}
	if spaces != 2 {
		errList = append(errList, ErrLineNotIncludesTwoSpaces)
	}

	if len(errList) == 0 {
		return nil
	}
	return errList
}

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
