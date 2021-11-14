package main
import "fmt"

func main(){
	name := "Jhoan"
	age := 32

	//Printf
	fmt.Printf("%s is %d old age\n", name, age)  

	//Srintf
	message := fmt.Sprintf("%s is %d old age", name, age) 
	fmt.Println(message)

	//Data Types
	fmt.Printf("name is: %T\n", name)
	fmt.Printf("Age is: %T\n", age)
}
