package main

import "fmt"

func main() {
	// comma-ok idiom
	// there’s no Try Catch block in Go.

	// 1. test for error on a function return
	// result, err := isEven(2)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// } else {
	// 	fmt.Printf("Result: %v\n", result)
	// }

	if result, err := isEven(3); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %v\n", result)
	}

	// 2. testing if a key-value item exists in a Go map
	person := map[string]string{
		"name": "John",
		"age":  "30",
		// "city": "London",
	}
	name, ok := person["name"]
	if ok {
		fmt.Printf("name: %v\n", name)
	} else {
		fmt.Println("name does not exist")
	}

	city, ok := person["city"]
	if ok {
		fmt.Printf("city: %v\n", city)
	}

	// 3. testing if an interface variable is of a certain Type
	var x any = "Hello, World!"
	str, ok := x.(string)
	if ok {
		fmt.Printf("x is a string: %v\n", str)
	} else {
		fmt.Println("x is not a string")
	}

	num, ok := x.(int)
	if ok {
		fmt.Printf("x is an int: %v\n", num)
	} else {
		fmt.Println("x is not an int")
	}

	// 4. testing if a channel is closed
	ch := make(chan int) // channel soon!
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for {
		if num, ok := <-ch; ok {
			fmt.Printf("Received: %d\n", num)
		} else {
			fmt.Println("Channel is closed")
			break
		}
	}
}

func isEven(num int) (bool, error) {
	if num%2 == 0 {
		return true, nil
	}
	return false, fmt.Errorf("%d is not an even number", num)
}
