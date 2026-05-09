package main

import "fmt"

func main() {

	// for loop
	for i := 0 ; i < 11 ; i++ {
		fmt.Printf("that is %d \n", i)
	}

	fmt.Printf("-------\n")

	// continue - it will jump a step 
	for c := 0 ; c < 5 ; c++ {
		if c == 3 {
			continue
		}
		fmt.Println(c)
	}

	fmt.Printf("-------\n")

	// break - it will stop for loop
	for n := 0 ; n < 5 ; n++ {
		if n == 3 {
			break
		}
		fmt.Println(n)
	}

	fmt.Printf("-------\n")

	// nested loop
	var arr = [2][]int{{1,2},{3,4,5}}

	for e := 0 ; e < len(arr) ; e++ {
		for d := 0 ; d < len(arr[e]) ; d++ {
			fmt.Println(arr[e][d])
		}
	}

	fmt.Printf("-------\n")

	// with range index and value
	// to ignore one of them use _
	for i, v := range arr {

		fmt.Println(i,v)

	}

}