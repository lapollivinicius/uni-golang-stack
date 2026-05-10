package main

import (
	"html/template"
	"os"
)

type person struct {
	Name string
	Age uint
}

func main() {

	// t, err := template.New("test").Parse("<h1> {{ . }} </h1>\n")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	data := person{Name: "john", Age: 20}
	// use Req instead os to response
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
