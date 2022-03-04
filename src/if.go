package main

import (
	"fmt"
	"log"
	"strconv"
)

func isPairNumber(num int) bool {
	return num%2 == 0
}

func isMachtUser(user, password, userDB, passwordDB string) bool {
	return user == userDB && password == passwordDB
}

func main() {
	values := []int{123, 453}

	if values[0] == 123 {
		fmt.Println("Is correct")
	} else {
		fmt.Println("Not is correct")
	}

	// AND
	if values[0] == 123 && values[1] == 453 {
		fmt.Println("Corrent AND")
	} else {
		fmt.Println("Not corrent AND")
	}

	// OR
	if values[0] == 123 || values[1] == 45344 {
		fmt.Println("Corrent OR")
	} else {
		fmt.Println("Not corrent OR")
	}

	// Error Controls
	num, err := strconv.Atoi("232")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nValue: %d", num)

	userFound := isMachtUser("Kratos", "12345678", "Kratos", "12345678")

	if !userFound {
		fmt.Println("\nUser not found")
	}

	fmt.Printf("\nUser found: %t", userFound)

	number := 4
	isPairNumber := isPairNumber(number)

	if !isPairNumber {
		fmt.Printf("\nThe number %d is not pair", number)
	}

	fmt.Printf("\nThe number %d is pair", number)
}
