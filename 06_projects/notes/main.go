package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type note struct {
	Title string
	Content string
}

func home(w http.ResponseWriter, r *http.Request) {

	// getting from URL
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	// templates dir
	files := []string{
		"view/template/home.html",
		"view/template/base.html",
	}
	t, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// note exemple
	data := note{
		Title: "My First Note",
		Content: "Lorem ipsum dolor, sit amet consectetur adipisicing elit.", 
	}
	t.ExecuteTemplate(w, "base", data)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing note id", http.StatusNotFound)
		return
	}

	files := []string{
		"view/template/base.html",
		"view/template/note.html",
	}

	t, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"view/template/base.html",
		"view/template/create.html",
	}
	w.Header().Add("Content-Type", "text/html")
	t, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func main() {

	fmt.Println("server running at 127.0.0.1:3000")

	// create server mux
	mux := http.NewServeMux()

	// routing static elements
	staticHandle := http.FileServer(http.Dir("view/static/"))

	// route with strip (cut the prefix /view/)
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandle))
	
	// routing and handling
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	// listener server
	http.ListenAndServe(":3000", mux)

}
