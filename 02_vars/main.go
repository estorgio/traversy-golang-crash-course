package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int, int8, int32, int64
	// uint, uint8, uint32, uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32, float64
	// complex64, complex128

	// Using var
	// var name = "Brad"
	var age int32 = 37
	var isCool = true
	// var size float32 = 2.3
	isCool = false

	// Shorthand
	// name := "Brad"
	// size := 1.3

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	// fmt.Println(name, email)
	fmt.Printf("%T\n", isCool)
}
