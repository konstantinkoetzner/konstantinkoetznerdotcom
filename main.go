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
	mux.HandleFunc("/font/", fontHandler)
	http.ListenAndServe(":8080", mux)
	fmt.Println("Server is running on port 8080")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	err := t.ExecuteTemplate(w, "index.gohtml", nil)
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

func fontHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("font", filepath.Clean(r.URL.Path[len("/font/"):]))
	http.ServeFile(w, r, filePath)
}
