package main

import (
	"fmt"
	"sort"
)

func main() {
	// arrays are fixed size, they cannot grow or shrink in size
	var number [5]int // array of 5 integers
	number[0] = 10
	number[1] = 20
	number[2] = 30
	number[3] = 40
	number[4] = 50

	fmt.Println(number)
	fmt.Println(number[0], number[1], number[2], number[3], number[4])
	fmt.Println(len(number))

	// array slicing
	// array[start:end]
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr[0:2]) // [1 2]
	fmt.Println(arr[:2]) // [1 2]
	fmt.Println(arr[2:]) // [3 4]
	fmt.Println(arr[:]) // [1 2 3 4]

	// sorting array
	arr2 := []int{5,3,2,4,1}
	fmt.Println(arr2) // before sorting
	sort.Ints(arr2[:])
	fmt.Println(arr2) // after sorting

	// array copying
	arr3 := []int{1,2,3,4,5}
	fmt.Println(arr3)
	arr4 := make([]int, len(arr3)) // make a new array of the same length as arr3
	copy(arr4,arr3) // copy the contents of arr3 into arr4
	fmt.Println(arr4)

	// reverse array
	arr5 := []int{1,2,3,4,5}
	fmt.Println(arr5)
	sort.Sort(sort.Reverse(sort.IntSlice(arr5[:])))
	fmt.Println(arr5)
};