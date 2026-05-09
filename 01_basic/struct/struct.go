package main

import "fmt"

// type name struct { name datatype }
type person struct {
	name string
	age int
	height float32
}

func main() {

	// instance
	var p1 person

	// assign values
	p1.name = "John"
	p1.age = 30
	p1.height = 1.84

	fmt.Println("name:", p1.name)
	fmt.Println("age:", p1.age)
	fmt.Printf("height: %.2fm \n", p1.height)

}