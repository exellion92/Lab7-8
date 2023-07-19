package main

import (
	"html/template"
	"net/http"
)

type HeaderData struct {
	Page string
}

type FooterData struct {
	URL  string
	Name string
	Year int
}

type PageData struct {
	HeaderData HeaderData
	FooterData FooterData
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/contact/", contactHandler)
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, templateFiles []string, data interface{}) {
	tmpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"templates/home.html",
		"templates/header.html",
		"templates/main.html",
		"templates/footer.html",
	}

	headerData := HeaderData{
		Page: "/",
	}

	footerData := FooterData{
		URL:  "https://t.me/NOD_Reaper",
		Name: "Дмитро",
		Year: 2023,
	}

	pageData := PageData{
		HeaderData: headerData,
		FooterData: footerData,
	}

	renderTemplate(w, templateFiles, pageData)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"templates/about.html",
		"templates/header.html",
		"templates/footer.html",
	}

	headerData := HeaderData{
		Page: "about",
	}

	footerData := FooterData{
		URL:  "https://t.me/NOD_Reaper",
		Name: "Дмитро",
		Year: 2023,
	}

	pageData := PageData{
		HeaderData: headerData,
		FooterData: footerData,
	}

	renderTemplate(w, templateFiles, pageData)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	templateFiles := []string{
		"templates/contact.html",
		"templates/header.html",
		"templates/footer.html",
	}

	headerData := HeaderData{
		Page: "contact",
	}

	footerData := FooterData{
		URL:  "https://t.me/NOD_Reaper",
		Name: "Дмитро",
		Year: 2023,
	}

	pageData := PageData{
		HeaderData: headerData,
		FooterData: footerData,
	}

	renderTemplate(w, templateFiles, pageData)
}
