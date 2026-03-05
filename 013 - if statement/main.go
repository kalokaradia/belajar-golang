package main

import "fmt"

func main() {
	x := 10

	// if statement
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 { // else if statement
		fmt.Println("x is less than 5")
	} else { // else statement
		fmt.Println("x is equal to 5")
	}
}
