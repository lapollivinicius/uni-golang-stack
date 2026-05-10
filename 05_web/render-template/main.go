package main

import (
	"fmt"
	"net/http"
)

// returning a simple html element
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1> Home </h1>")

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/home", Home)
	
	http.ListenAndServe(":3000", mux)

}