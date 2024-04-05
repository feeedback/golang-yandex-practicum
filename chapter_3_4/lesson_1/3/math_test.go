package math // Тема 4/6: Тестирование → Урок 1/3 "Тестирование"
// https://practicum.yandex.ru/trainer/go-basics/lesson/d01326ef-d2a3-46b1-8dd1-f4a7f6fa8d0e/

import "testing"

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}

	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}

	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}

	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 2)
	if err == nil {
		t.Error("first arg is zero  - expected error not be nil")
	}

	_, err = Add(1, 0)
	if err == nil {
		t.Error("second arg is zero  - expected error not be nil")
	}

	_, err = Add(0, 0)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}
