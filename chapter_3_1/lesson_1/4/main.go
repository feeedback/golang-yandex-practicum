package main // Тема 1/6: ООП → Урок 1/3 "Методы"
// https://practicum.yandex.ru/trainer/go-basics/lesson/59203fa8-851b-444a-82f5-d031138079d2/

import (
	"fmt"
	"time"
)

// Реализуйте тип Stopwatch, который будет сохранять время каждой промежуточной фиксации секундомера и выдавать результаты относительно общего времени старта.
// Тип должен обладать следующими методами:
// Start() — запустить/сбросить секундомер;
// SaveSplit() — сохранить промежуточное время;
// GetResults() []time.Duration — вернуть текущие результаты.

type Stopwatch struct {
	startTime time.Time
	splits    []time.Time
}

func (sw *Stopwatch) Start() {
	sw.startTime = time.Now()
	sw.splits = nil
}

func (sw *Stopwatch) SaveSplit() {
	sw.splits = append(sw.splits, time.Now())
}

func (sw Stopwatch) GetResults() (retResults []time.Duration) {
	for _, splitTime := range sw.splits {
		retResults = append(retResults, splitTime.Sub(sw.startTime))
	}

	return
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
