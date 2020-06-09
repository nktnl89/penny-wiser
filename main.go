package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	// items := []Item{Item{0, "Кредит"}, Item{1, "Кварплата"}}

	invoices := []Invoice{Invoice{0, "Альфа", 0}, Invoice{1, "Тинькоф", 0}, Invoice{2, "Нал", 0}, Invoice{3, "Черный День", 150000}}

	err := tpl.ExecuteTemplate(res, "invoices.gohtml", invoices)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
