package main
import "fmt"
import "math"

func main() {
	b := 10
	h := 20

	rectangle := b * h
	fmt.Println("Rectangle area", rectangle)
	
	B := 7
	trapeze := (h * ( B * h ))/2
	fmt.Println("Trapeze area:", trapeze)

	var r float64 = 50
	cicle := math.Pi * math.Pow(r, 2)
	fmt.Print("Cicle area:", math.Round(cicle))
}
