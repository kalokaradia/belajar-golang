package main

import "fmt"

func main() {
	// creating slices
	// 1. using slice literal
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	// 2. using make function
	slice2 := make([]int, 5)     // Length 5, capacity 5
	slice3 := make([]int, 5, 10) // Length 5, capacity 10
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(len(slice2))
	fmt.Println(len(slice3))
	fmt.Println(cap(slice2))
	fmt.Println(cap(slice3))
	// 3. from existing array
	arr := [5]int{1, 2, 3, 4, 5}
	slice4 := arr[1:4]
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// slice operations
	// 1. appending to slice
	slice5 := []int{1, 2, 3}
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	slice5 = append(slice5, 4, 5)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	// 2. copying slices
	slice6 := make([]int, 5)
	fmt.Println(slice6)
	// fmt.Println(len(slice6))
	// fmt.Println(cap(slice6))
	copy(slice6, slice5)
	fmt.Println(slice6)
	// fmt.Println(len(slice6))
	// fmt.Println(cap(slice6))

	// slice manipulation
	// 1. slicing a slice
	/*
		slice[start:end]   // Elements from start to end-1
		slice[start:]      // Elements from start to end
		slice[:end]        // Elements from beginning to end-1
		slice[:]           // Entire slice
	*/
	slice7 := slice5[1:4] // index 1 to
	fmt.Println(slice7)
}
