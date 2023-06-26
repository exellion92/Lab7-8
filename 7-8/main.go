package main

import (
	"html/template"
	"net/http"
)

type Developer struct {
	Name         string
	Organization string
	Year         int
	Url          string
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")

	developer := Developer{"Дмитро", "НУВГП", 2023, "https://t.me/NOD_Reaper"}

	tmpl.Execute(w, developer)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/about.html")

	developer := Developer{"Дмитро", "НУВГП", 2023, "https://t.me/NOD_Reaper"}

	tmpl.Execute(w, developer)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8080", nil)
}
