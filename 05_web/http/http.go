package main

import (
	"fmt"
	"net/http"
)

// struct handle
type MyHandler struct{}

// function that will handle 
func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// header - return from server to cliente (response header)
	w.WriteHeader(201)
	w.Write([]byte("Hello!"))
		
}

// function without struct and ServerHTTP
func SayHi(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello"))

}

// take data from URL
func byUrl(w http.ResponseWriter, r *http.Request) {

	// recovery data using URL
	urlPath := r.URL.Path
	urlQuery := r.URL.Query()

	fmt.Fprint(w, urlPath)
	fmt.Fprint(w, urlQuery)

}

// methods
func handleMethods(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		// to pass the status code
		w.WriteHeader(405)

		// but it will be pass with error
		// http.Error(w, "method not allowed", 405)
		return
	}

	// to set something at header (allow a specific method)
	w.Header().Set("Allow", http.MethodGet)
	w.Header().Set("Content-Type", "application/json")

	// to "delete" something in header
	w.Header()["Date"] = nil

	fmt.Fprint(w, `{ "name": "john" }`)

}

func main() {

	fmt.Println("server running at port 127.0.0.1:3000")

	// handler instance
	handler := MyHandler{}

	// pattern 
	http.Handle("/", handler)

	// pattern with handle func (handler func is a struct and a func)
	http.DefaultServeMux.HandleFunc("/hello", SayHi)

	// we create ourselves server mux (mix routes)
	// or we routing using mux and use mux to routin
	// or we use nil and by default routing with http...
	mux := http.NewServeMux()
	mux.HandleFunc("/ola", SayHi)

	// using a annon func and Fprint that convert bytes to strings
	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "You are in our product page!")
	})

	// working with methods
	mux.HandleFunc("/methods", handleMethods)

	// recovery data URL
	mux.HandleFunc("/test", byUrl)

	// funcion that listen a port to server something
	// when we use nil, it will run ServerMux by default
	// it means that its will routing with handlers
	// can be a handler func (nil)
	http.ListenAndServe("127.0.0.1:3000", mux) 

}