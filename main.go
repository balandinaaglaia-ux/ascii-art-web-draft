package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Input  string
	Banner string
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	renderPage(w, PageData{Banner: "standard"})
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10000)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	input := r.FormValue("text")
	bannerName := r.FormValue("banner")

	if input == "" {
		http.Error(w, "Text cannot be empty", http.StatusBadRequest)
		return
	}

	input = convertNewLines(input)

	for _, char := range input {
		if char != '\n' && (char < 32 || char > 126) {
			http.Error(w, "Invalid character", http.StatusBadRequest)
			return
		}
	}

	if bannerName != "standard" &&
		bannerName != "shadow" &&
		bannerName != "thinkertoy" {
		http.Error(w, "Invalid banner", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile("banners/" + bannerName + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Banner not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	result, err := printAscii(input, string(data))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	renderPage(w, PageData{
		Input:  input,
		Banner: bannerName,
		Result: result,
	})
}

func renderPage(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Template not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	log.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//    http://localhost:8080
