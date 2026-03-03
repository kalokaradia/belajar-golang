package main

import "fmt"

func main () {
	// NUMERIC TYPES
	// signed integer (value can negative or positive)
	/*
		int
		int8
		int16
		int32
		int64
	*/
	var age int8 = 20
	var poin int8 = -25
	
	fmt.Printf("%v, %v\n", age, poin)
	
  // unsigned integer (value can't negative)
	/*
	uint
	uint8
	uint16
	uint32
	uint64
	*/
	var score uint8 = 20
	// var cat uint8 = -5 -> error
	var cat uint8 = 5
	
	fmt.Printf("%v, %v\n", score, cat)
	
	// floating points
	// float32 -> ~6–7 digit decimal
	// float64 -> ~15–16 digit decimal
	var height float32 = 1.6
	fmt.Printf("%v\n", height)
	
	// complex number (it doesn't really matter to me now)
	
	// BOOLEAN
	var isGraduate bool = true
	var isAndroidProgrammer bool = false
	fmt.Printf("%v, %v\n", isGraduate, isAndroidProgrammer)
	
	// RUNE
	// rune = alias for int32
	// runes use single quotes to write their values.
	var smileEmoji rune = '🙂'
	var r int32= '你'
	fmt.Printf("%c, %c\n", smileEmoji,r)
	
	// STRING
	var name string = "Kaloka Radia Nanda"
	var countryName string = "Indonesia"
	fmt.Printf("%v, %v\n", name, countryName)
}