package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/template/home.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == " " {
		http.Error(w, "missing note id", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("view/template/note.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create a Note")
}



func main() {

	fmt.Println("server running at 127.0.0.1:3000")

	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":3000", mux)


}
