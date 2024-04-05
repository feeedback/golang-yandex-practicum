package main // Тема 4/6: Тестирование → Урок 2/3 "Интерфейсы в тестировании"
// https://practicum.yandex.ru/trainer/go-basics/lesson/772e29d4-65f1-40cd-8f7c-253ccdbf13a7/

type APIClient interface {
	GetData(query string) (Response, error)
}

type Response struct {
	Text       string
	StatusCode int
}

type MockAPIClient interface {
	APIClient
	SetResponse(resp Response, err error)
}

type Mock struct {
	Resp Response
	Err  error
}

func (m *Mock) GetData(query string) (Response, error) {
	return m.Resp, m.Err
}

func (m *Mock) SetResponse(resp Response, err error) {
	m.Resp = resp
	m.Err = err
}
