package main

import "fmt"

func main() {

	msg := SayHello("John")

	fmt.Println(msg)

	Recursion(5)

}

// basic structure
// func Name ( parameters type ) type output { return value }
func SayHello(name string) string {
	return "Hi, " + name
}

// returns and multiple
func SumAndTimes(x,y int) (sum int, times int) {
	sum = x + y
	times = x * y
	return
}

// recursion
func Recursion(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Println(n)
	return Recursion(n - 1)
}