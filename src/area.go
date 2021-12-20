package main

import (
	"fmt"
	"math"
)

func main() {
	base := 10
	height := 20

	rectangle := rectangleAreaCalculate(base, height)
	fmt.Println("Rectangle area", rectangle)

	base2 := 7
	trapeze := trapezeAreaCalculate(base, base2, height)
	fmt.Println("Trapeze area:", trapeze)

	var radius float64 = 50
	cicle := cicleRadiusCalculate(radius)
	fmt.Print("Cicle area: ", math.Round(cicle))
}

func rectangleAreaCalculate(base int, height int) int {
	return base * height
}

func trapezeAreaCalculate(base, base2, height int) int {
	return height * (base + base2) / 2
}

func cicleRadiusCalculate(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}
