package main

import "fmt"

func main() {

	a := 10

	// single case
	switch a {
		case 10:
			fmt.Println("it's ten")
		case 20:
			fmt.Println("it's twenty")
		default:
			fmt.Println("everything else")
	}

	// multicases
	switch a {
		case 10, 20:
			fmt.Println("it's ten or twenty")
	}




}