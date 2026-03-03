package main

import "fmt"

func main() {
	x := 42

	if true {
		x:= 100
		fmt.Println(x) // This will print 100, because the x variable declared inside the if block shadows the outer x variable.
	}

	fmt.Println(x) // This will print 42, because the x variable declared inside the if block does not affect the outer x variable.
}