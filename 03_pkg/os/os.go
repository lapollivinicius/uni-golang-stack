package main

import (
	"fmt"
	"os"
)

func main() {

	// OS WITH ENV ------------>

	// to get a env value from system envs
	port := os.Getenv("SERVER_PORT")
	fmt.Println(port)

	// to check if an env exists use lookupenv
	port1, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		panic(ok)
	}
	fmt.Println(port1, ok)


}
