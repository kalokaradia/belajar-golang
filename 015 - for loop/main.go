package main

import "fmt"

func main() {
	/*
		* for loop syntax
		for initialisation; condition; post {
			code
		}
	*/

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}

	// break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nloop ended\n")

	// continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nloop ended\n")

	// nested for loops
	n := 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	// while loop using for loop
	x := 5
	for x <= 10 {
		fmt.Printf("%d ", x)
		x++
	}
	fmt.Printf("\n")

	// infinite loop
	// for { fmt.Printf("Infinity Loop")	}
}
