package main

import "fmt"

func main() {
	day := 9

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	/*
		case 6:
			fmt.Println("Saturday")
			* Duplicate cases are not allowed
	*/
	case 7:
		fmt.Println("Sunday")
	default: // default case if no other case matches
		fmt.Println("Invalid day")
	}

	// multiple expressions in case
	letter := "b"
	switch letter {
	case "a", "i", "u", "e", "o": // multiple cases can be combined
		fmt.Printf("%s is a vowel\n", letter)
	default:
		fmt.Printf("%s is not a vowel\n", letter)
	}

	// expressionless switch
	hour := 15 // hour in 24 hour format
	// using switch to determine the work shift
	switch {
	case hour >= 6 && hour < 12:
		fmt.Println("It's the morning shift.")
	case hour >= 12 && hour < 17:
		fmt.Println("It's the afternoon shift.")
	case hour >= 17 && hour < 21:
		fmt.Println("It's the evening shift.")
	case (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6):
		fmt.Println("It's the night shift.")
	default:
		fmt.Println("Invalid hour.")
	}

	// fallthrough
	// * fallthrough happens even when the case evaluates to false
	score := 85

	switch {
	case score >= 90:
		fmt.Println("A")
		fallthrough
	case score >= 80:
		fmt.Println("better than 80")
	}
}
