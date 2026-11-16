package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	renderPage(w, PageData{})
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit request size
	r.Body = http.MaxBytesReader(w, r.Body, 10000)

	// Read form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	input := r.FormValue("text")
	bannerName := r.FormValue("banner")

	// Check empty input
	if input == "" {
		http.Error(w, "Text cannot be empty", http.StatusBadRequest)
		return
	}

	// Convert escaped newlines
	input = convertNewLines(input)

	// Validate characters
	for _, char := range input {
		if char != '\n' && (char < 32 || char > 126) {
			http.Error(w, "Invalid character", http.StatusBadRequest)
			return
		}
	}

	// Select banner
	if bannerName != "standard" &&
		bannerName != "shadow" &&
		bannerName != "thinkertoy" {

		http.Error(w, "Invalid banner", http.StatusBadRequest)
		return
	}

	// Read banner file
	bannerPath := "banners/" + bannerName + ".txt"

	data, err := os.ReadFile(bannerPath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Banner not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	banner := string(data)

	// Generate ASCII art
	result := printAscii(input, banner)

	// Display result
	renderPage(w, PageData{
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

	// Execute template into a buffer first
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

	// 2 next are for letting go interact with files from the static folder
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