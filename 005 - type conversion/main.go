package main

import "fmt"

func main() {
	// The expression T(v) converts the value v to the type T.
	var score int8 = 100
	var newScore int16 = int16(score)

	fmt.Printf("%v, %v\n", score, newScore)
	fmt.Printf("%T, %T\n", score, newScore)
}
