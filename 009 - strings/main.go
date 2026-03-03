package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// string literals
	message := "Hello, World!"
	fmt.Println(message)
	// raw string literals
	rawMessage := `This is a raw string literal.
It can span multiple lines and includes "quotes" without needing to escape them.`
	fmt.Println(rawMessage)
	// string concatenation
	greeting := "Hello " + "Go!"
	fmt.Println(greeting)
	// string conversions
	number := 42
	numberStr := fmt.Sprintf("%d", number)
	fmt.Printf("%T %T\n", number, numberStr)
	// calculating string length
	length := len(message)
	fmt.Printf("Length of message: %d\n", length)
	text := "Hello, 你好"
	fmt.Printf("Length of text isn't %d\n", len(text))
	fmt.Printf("Length of text is %d\n", utf8.RuneCountInString(text))
}
