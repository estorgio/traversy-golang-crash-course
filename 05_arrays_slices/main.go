package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitArr := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArr)

	// Count array length
	fmt.Println(len(fruitArr))

	// return array range
	fmt.Println(fruitArr[1:3])
}
