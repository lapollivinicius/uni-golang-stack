package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	// condition if
	if a > b {
		fmt.Println(a)
	}

	// if and else
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	// if and else if
	if a > b {
		fmt.Println(a)
	} else if a == b {
		fmt.Println(b)
	}

	// nested condition
	if a < b {
		if a != b {
			fmt.Println(b)
		}
	}

}