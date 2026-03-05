package main

import "fmt"

func main() {
	// for range
	/*
		* for range syntax
		for index, value := range data {
			code
		}
	*/

	nums := [5]string{"h", "e", "l", "l", "o"}
	for i, v := range nums {
		fmt.Printf("%v. %v\n", i+1, v)
	}
	// without index
	for _, v := range nums {
		fmt.Printf("%v", v)
	}
	fmt.Printf("\n")

	// map iteration
	person := map[string]any{
		"name":     "Kaloka Radia Nanda",
		"age":      20,
		"birthday": "1 January 1000",
		"address":  "123 Street won't forget",
	}

	for k, v := range person {
		fmt.Printf("%v : %v\n", k, v)
	}

	// string iteration
	msg := "hello"
	for _, v := range msg {
		fmt.Printf("%c ", v)
	}
}
