package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/language", languageHandler)
	mux.HandleFunc("/css/", cssHandler)
	mux.HandleFunc("/img/", imgHandler)
	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lang := determineLanguage(r)
	err = t.Execute(w, lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func determineLanguage(r *http.Request) string {
	c, err := r.Cookie("language")
	if err != nil {
		return "en"
	}
	if c.Value == "de" {
		return "de"
	}
	return "en"
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	lang := r.FormValue("lang")
	if lang != "en" && lang != "de" {
		http.Error(w, "Invalid language", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "language",
		Value:    lang,
		Path:     "/",
		MaxAge:   int(365 * 24 * time.Hour / time.Second),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("css", filepath.Clean(r.URL.Path[len("/css/"):]))
	http.ServeFile(w, r, filePath)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("img", filepath.Clean(r.URL.Path[len("/img/"):]))
	http.ServeFile(w, r, filePath)
}
