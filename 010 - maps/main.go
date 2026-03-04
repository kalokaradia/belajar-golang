package main

import "fmt"

func main() {
	// create a map
	// 1. using make function
	person := make(map[string]string)
	person["name"] = "John"
	person["age"] = "30"
	fmt.Printf("name: %v\n", person["name"])
	fmt.Printf("age: %v\n\n", person["age"])
	// 2. using map literal
	person2 := map[string]string{
		"name": "Jane",
		"age":  "25",
	}
	fmt.Printf("name: %v\n", person2["name"])
	fmt.Printf("age: %v\n", person2["age"])

	// operations with maps
	// 1. delete a key-value pair
	delete(person, "age")
	fmt.Printf("person after delete: %v\n", person)
	// 2. check if a key exists
	name, ok := person["name"]
	if ok { // conditional statement soon
		fmt.Printf("name: %v\n", name)
	} else {
		fmt.Println("name does not exist")
	}
	// 3. iterate over a map
	fmt.Println("iterating over person2:")
	for key, value := range person2 { // range statement soon
		fmt.Printf("%s: %s\n", key, value)
	}
	// 4. adding a new key-value pair
	person2["city"] = "New York"
	fmt.Printf("person2 after adding city: %v\n", person2)
}
