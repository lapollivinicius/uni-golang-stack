package main

import "fmt"

func main() {

	// Arithmetic
	add := 10 + 20

	sub := 10 - 20

	div := 20 / 10

	mul := 10 * 20

	mod := 10 % 20

	// increment and decrement
	add++
	sub--

	// assignment
	var a int = 10

	a += 10 // 11
	a -= 10 // 0
	a *= 10 // 100
	a /= 10 // 1
	a %= 10 // 0

	// comparison 
	var b int = 20
	var result bool

	result = a == b
	result = a != b
	result = a > b
	result = a < b
	result = a >= b
	result = a <= b

	// logical
	result = a < b && b > a
	result = a < b || b > a
	result = !result

	fmt.Println(add, sub, div, mul, mod, result)

}