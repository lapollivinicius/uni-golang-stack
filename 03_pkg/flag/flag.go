package main

import (
	"flag"
	"fmt"
)

func main() {

	// flag string - define name and value to use on cmd line
	port := flag.String("port", "3000", "server port")

	// with stringvar
	var host string
	flag.StringVar(&host, "host", "http://localhost:", "server host")

	// boolean - if it has (name) is true if it hasnt is false
	var turbomode bool
	flag.BoolVar(&turbomode, "turbo", false, "turbo mode") 

	// go on the flags to be used
	flag.Parse()

	// that is a ponteiro
	fmt.Println(*port)

	// check boolean flag
	if turbomode { fmt.Print("turbomode on ") }

	// that is a simple var
	fmt.Print(host)
	fmt.Print(*port)

	fmt.Print("\n")

	// use prefix (-) and (=) to define the value
	// ex. -port=3000
	// use (-h) help to know the cmd parameters
}	