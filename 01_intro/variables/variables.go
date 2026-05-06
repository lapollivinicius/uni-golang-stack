package main

import "fmt"

func main() {

	// var name type = "value"

	// let start with types

	var Inteiger int = 100 // range -123 to 123

	var String string = "hello" 

	var FloatandDouble float32 = 19.12 // int32 float32 define the bits

	var Boolean bool = false 

	// you need declare type or value
	// the type of the variable is inferred with :=
	variable := "string"
	

	fmt.Println(variable)

	// multiple 
	var a, b, c int = 1, 2, 3

	// different type
	s, i := "string", 100

	// in a block
	var (
		d int
		e string = "hello"
	)

	fmt.Println(Inteiger, String, FloatandDouble, Boolean, a, b, c, d, e, s, i)

	// NAME RULE
	// start with only _ or letter
	// can contain only alphanumerics and (_)
	// cannot contain spaces or be a keyword

	// recomend camelCase, PascalCase or snake_case

	// constants unchangeable, read-only
	const read string = "i'm not a changing text"

	// also can be in block and untyped
	const (
		onlyToSee = "yes, baby"
	)
}
	



