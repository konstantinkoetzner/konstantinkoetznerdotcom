package main

import (
	"fmt"
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
	mux.HandleFunc("/person", personHandler)
	mux.HandleFunc("/publication", publicationHandler)
	mux.HandleFunc("/contact", contactHandler)
	mux.HandleFunc("/veroeffentlichung", veroeffentlichungHandler)
	mux.HandleFunc("/kontakt", kontaktHandler)
	http.ListenAndServe(":8080", mux)
	fmt.Println("Server is running on port 8080")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	lang := determineLanguage(r)
	err := t.ExecuteTemplate(w, "index.gohtml", lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	lang := determineLanguage(r)
	err := t.ExecuteTemplate(w, "person.gohtml", lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func publicationHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	lang := determineLanguage(r)
	err := t.ExecuteTemplate(w, "publication.gohtml", lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	lang := determineLanguage(r)
	err := t.ExecuteTemplate(w, "contact.gohtml", lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func veroeffentlichungHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	lang := determineLanguage(r)
	err := t.ExecuteTemplate(w, "veroeffentlichung.gohtml", lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func kontaktHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	lang := determineLanguage(r)
	err := t.ExecuteTemplate(w, "kontakt.gohtml", lang)
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
	ref := r.Header.Get("Referer")
	if r.Method != http.MethodPost {
		if ref == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, ref, http.StatusSeeOther)
		return
	}
	uu := []string{
		"publication", "veroeffentlichung",
		"contact", "kontakt",
	}
	lang := r.FormValue("lang")
	if lang != "en" && lang != "de" {
		http.Error(w, "Invalid language", http.StatusBadRequest)
		return
	}
	for i, u := range uu {
		if u == filepath.Base(ref) {
			if i%2 == 0 {
				setLanguageCookie(w, lang)
				http.Redirect(w, r, fmt.Sprintf("/%s", uu[i+1]), http.StatusSeeOther)
				return
			}
			if i%2 == 1 {
				setLanguageCookie(w, lang)
				http.Redirect(w, r, fmt.Sprintf("/%s", uu[i-1]), http.StatusSeeOther)
				return
			}
		}
	}
	setLanguageCookie(w, lang)
	if ref == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, ref, http.StatusSeeOther)
}

func setLanguageCookie(w http.ResponseWriter, lang string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "language",
		Value:    lang,
		Path:     "/",
		MaxAge:   int(365 * 24 * time.Hour / time.Second),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("css", filepath.Clean(r.URL.Path[len("/css/"):]))
	http.ServeFile(w, r, filePath)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("img", filepath.Clean(r.URL.Path[len("/img/"):]))
	http.ServeFile(w, r, filePath)
}
