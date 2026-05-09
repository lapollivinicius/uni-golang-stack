package main

import "fmt"

func main() {

	// var arrayName = [ len ] Type { values }
	var letters = [10]string{"a", "b", "c", "d", "e"}

	// they are indexed and type unchangeable

	// [...] inferred length
	var numbers = [...]int{1, 2, 3, 4, 5}

	// to index use [position] 
	a := letters[0]

	// and assign with {index:value}
	fruits := [3]string{1:"apple", 2:"pineapple"}

	// length - use len() to know the capacitable of an array
	length := len(letters) 

	fmt.Println(letters[2], numbers[2], fruits[2], a, length)


}
