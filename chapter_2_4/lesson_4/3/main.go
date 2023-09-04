package main // Тема 1/4: Композитные типы → Урок 4/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/b1ee4a40-18d6-44f2-ba62-b2918538f811/

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	person := Person{
		Name:        "Aлекс",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	jsonData, _ := json.Marshal(person)

	fmt.Println(string(jsonData))
}
