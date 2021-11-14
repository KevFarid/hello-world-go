package main
import "fmt"

func main(){
	fmt.Println("Hello world")

	number1 := 31
	var number2 int = 78
	var result int

	fmt.Println(number1, number2, result)

	//Sum
	result = number1 + number2
	fmt.Println("Sum:", result)

	//rest
	result = number1 - number2
	fmt.Println("Rest:", result)

	//multi
	result = number1 * number2
	fmt.Println("Multi:", result)

	//div
	result = number1 / number2
	fmt.Println("div:", result)

	number := 20

	//incremental
	number++
	fmt.Println("Inclemental:", number)

	//Decremental
	number--
	fmt.Println("Decremental:", number)
}
