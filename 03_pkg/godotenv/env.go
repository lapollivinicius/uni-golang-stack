package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
) 

func main() {
	
	//init
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Printf("https://%s:%s", host, port)

}