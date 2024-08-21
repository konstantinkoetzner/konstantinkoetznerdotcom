package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/css/", cssHandler)
	mux.HandleFunc("/img/", imgHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("css", filepath.Clean(r.URL.Path[len("/css/"):]))
	http.ServeFile(w, r, filePath)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("img", filepath.Clean(r.URL.Path[len("/img/"):]))
	http.ServeFile(w, r, filePath)
}
