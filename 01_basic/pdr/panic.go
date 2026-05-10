package main

import "fmt"

func main() {

	// panic - stop flow and can throw something
	a := [4]int{1,2,3}

	for c := 0 ; c < 3 ; c++ {

		if c > 1 {
			panic(c)
		}
		fmt.Println("the number is ", a[c])

	} 
		defer err()
}

func err() {

	if r := recover(); r != nil {
		fmt.Println(r)
	}

}
