package main

import "fmt"

func main() {

	defer fmt.Println("DEFER")

	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 44, 30}

	fizzbuzz(numberList)

	findNumber := findIndexNumber(numberList, 30)

	fmt.Println(findNumber)

}

func fizzbuzz(numberList []int) {
	for _, number := range numberList {
		if number%3 == 0 && number%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if number%3 == 0 {
			fmt.Println("fizz")
			continue
		}
		if number%5 == 0 {
			fmt.Println("buzz")
			continue
		}
	}
}

func findIndexNumber(numberList []int, numberToFind int) int {
	index := 0
	finishArray := false
	for i, number := range numberList {
		if number == numberToFind {
			index = i
			break
		}
		if i == len(numberList)-1 {
			finishArray = true
		}
	}

	if index == 0 && finishArray {
		index = -1
	}
	return index
}
