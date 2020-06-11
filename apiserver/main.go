package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/invoices", invoices)
	http.HandleFunc("/items", items)
	http.HandleFunc("/items/add", itemsAdding)
	http.HandleFunc("/plans", plans)
	http.HandleFunc("/periods", periods)

	_ = http.ListenAndServe(":8080", nil)
}

func items(res http.ResponseWriter, req *http.Request) {
	items := []Item{Item{0, "Кредит"}, Item{1, "Кварплата"}}

	err := tpl.ExecuteTemplate(res, "items.gohtml", items)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func itemsAdding(res http.ResponseWriter, req *http.Request) {
	//items := []Item{Item{0, "Кредит"}, Item{1, "Кварплата"}}
	title := req.FormValue("title")
	fmt.Println(title)
	err := tpl.ExecuteTemplate(res, "items.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func invoices(res http.ResponseWriter, req *http.Request) {
	invoices := []Invoice{Invoice{0, "Альфа", 0}, Invoice{1, "Тинькоф", 0}, Invoice{2, "Нал", 0}, Invoice{3, "Черный День", 150000}}

	err := tpl.ExecuteTemplate(res, "invoices.gohtml", invoices)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func periods(res http.ResponseWriter, req *http.Request) {
	timeZone, _ := time.LoadLocation("Europe/Samara")
	//"Russia Time Zone 3":              {"+04", "+04"}
	periods := []Period{Period{time.Date(2020, 6, 10, 0, 0, 0, 0, timeZone),
		time.Date(2020, 6, 24, 0, 0, 0, 0, timeZone)}}

	err := tpl.ExecuteTemplate(res, "periods.gohtml", periods)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func plans(res http.ResponseWriter, req *http.Request) {
	plan := req.FormValue("plan_id")
	if plan == "" {
		log.Fatal("there is no plan id")
	}
	err := tpl.ExecuteTemplate(res, "invoices.gohtml", plan)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
	fmt.Println("some plans", plan)

}

// Invoice ...
type Invoice struct {
	ID    int
	Title string
	Plan  int
}

// HasPlan ...
func (i Invoice) HasPlan() bool {
	return i.Plan > 0
}

// GetCurrentSum ...
func (i Invoice) GetCurrentSum() int {
	return 100 // надо прикрутить расчет текущего значения по таблице с items
}

// Item ...
type Item struct {
	ID    int
	Title string
}

// Period ...
type Period struct {
	StartDate  time.Time
	FinishDate time.Time
}
