package main

import "fmt"

func main() {
	// Declaring Constants

	// untyped constant
	const Name = "Kaloka Radia Nanda"
	const Pi = 3.14
	fmt.Println(Name, Pi)
	// typed constant
	const GoldenRatio float32 = 1.618
	const Age int = 100
	fmt.Println(GoldenRatio, Age)

	// iota
	// is an automatic counter inside a const block.
	const (
		Monday int = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
