package main

import (
	"flag"
	"fmt"
)

func main() {

	// flag string - define name and value to use on cmd line
	port := flag.String("port", "3000", "Server")

	// with stringvar
	host := 

	// go on the flags to be used
	flag.Parse()

	// that is a ponteiro
	fmt.Println(*port)

	// that is a simple var
	fmt.Println(host)
}	