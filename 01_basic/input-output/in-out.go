package main

// use this lib to import methods to input and output
import "fmt"

func main() {

	var text = "hello"

	// root print
	fmt.Print("no break line")
	fmt.Println("with break line")
	fmt.Printf("with format / concat %v", text)

	// input 
	var nome string
    var idade int

    fmt.Scan(&nome, &idade)

    fmt.Println(nome, idade)
}