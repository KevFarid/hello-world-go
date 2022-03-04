package main

import "fmt"

func main() {

	numberList := []int{1, 2, 3, 4, 5}

	// normal for
	for i := 0; i < len(numberList); i++ {
		fmt.Printf("\nNumber: %d", numberList[i])
	}

	//for while
	count := 0
	for count <= 10 {
		fmt.Printf("\nCount: %d", count)
		count++
	}

	//for forever
	/**
	* !for {
			forever
	* !}
	*/

	//for range
	for i, num := range numberList {
		fmt.Printf("\nRange: %d, value: %d", i, num)
	}
}
