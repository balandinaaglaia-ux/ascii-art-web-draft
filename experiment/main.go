package main

import (
"html/template"
"net/http"
)

func main() {
http.HandleFunc("/", func(w http.ResponseWriter, 
r *http.Request) { // read and write when going to web page
tmpl, _ := template.ParseFiles("index.html")
tmpl.Execute(w, "Hello, ASCII!")
})

http.ListenAndServe(":8080", nil)
}