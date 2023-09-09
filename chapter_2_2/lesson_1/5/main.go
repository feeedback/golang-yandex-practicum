package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

type calcShapeArea func(a float64) float64

func area(f figures) (calcShapeArea, bool) {
	switch f {
	case square:
		return func(a float64) float64 { return math.Pow(a, 2) }, true
	case circle:
		return func(a float64) float64 { return math.Pi * math.Pow(a, 2) }, true
	case triangle:
		return func(a float64) float64 { return (math.Sqrt(3) / 4) * math.Pow(a, 2) }, true

	default:
		return nil, false
	}
}

func main() {
	aSide := 2.0

	calcAreaSquare, _ := area(square)
	areaSquare := calcAreaSquare(aSide) // 4

	calcAreaCircle, _ := area(circle)
	areaCircle := calcAreaCircle(aSide) // 4*Pi

	calcAreaTriangle, _ := area(triangle)
	areaTriangle := calcAreaTriangle(aSide) // math.Sqrt(3)

	fmt.Println("Площадь фигуры:", areaSquare, areaCircle, areaTriangle)
}
