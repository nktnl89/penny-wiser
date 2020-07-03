package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type server struct {
	router    *mux.Router
	logger    *logrus.Logger
	store     store.Store
	templates *template.Template
}

func newServer(store store.Store) *server {
	s := &server{
		router:    mux.NewRouter(),
		logger:    logrus.New(),
		store:     store,
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/dashboard", s.handleDashboard()).Methods("GET")

	s.router.HandleFunc("/invoices", s.handleInvoices()).Methods("GET")
	s.router.HandleFunc("/invoices/update", s.handleInvoicesUpdate()).Methods("GET")
	s.router.HandleFunc("/invoices/update/process", s.handleInvoicesUpdateProcess()).Methods("POST")
	s.router.HandleFunc("/invoices/delete", s.handleInvoicesDelete()).Methods("GET")

	s.router.HandleFunc("/items", s.handleItems()).Methods("GET")
	s.router.HandleFunc("/items/update", s.handleItemsUpdate()).Methods("GET")
	s.router.HandleFunc("/items/update/process", s.handleItemsUpdateProcess()).Methods("POST")
	s.router.HandleFunc("/items/delete", s.handleItemsDelete()).Methods("GET")

	s.router.HandleFunc("/plans", s.handlePlans()).Methods("GET")
	s.router.HandleFunc("/plans/update", s.handlePlansUpdate()).Methods("GET")
	s.router.HandleFunc("/plans/update/process", s.handlePlansUpdateProcess()).Methods("POST")
	s.router.HandleFunc("/add", s.handleAddItem()).Methods("POST")
}

func (s *server) handleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// тут берем ентриес всякие и выводим текущее состояние по планам
	}
}

func (s *server) handleAddItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		sbs := string(bs)
		fmt.Println("USERNAME: ", sbs)

		_, _ = w.Write([]byte(sbs))
	}
}

func (s *server) handlePlans() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		plans, err := s.store.Plan().FindAll()
		if err != nil {
			log.Fatal(err)
		}
		err = s.templates.ExecuteTemplate(w, "plans.gohtml", plans)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) handlePlansUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		startDate, _ := time.Parse(time.RFC3339, r.URL.Query().Get("start_date"))
		finishDate, _ := time.Parse(time.RFC3339, r.URL.Query().Get("finish_date"))
		sum, _ := strconv.Atoi(r.URL.Query().Get("sum"))
		closed, _ := strconv.ParseBool(r.URL.Query().Get("closed"))
		allItems, _ := s.store.Item().FindAll()
		p := &model.Plan{
			ID:         id,
			Sum:        sum,
			StartDate:  startDate,
			FinishDate: finishDate,
			Closed:     closed,
			Items:      []*model.Item{},
			AllItems:   allItems,
		}

		err := s.templates.ExecuteTemplate(w, "plan-form.gohtml", p)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) handlePlansUpdateProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		id, _ := strconv.Atoi(r.FormValue("id"))
		startDate, _ := time.Parse(time.RFC3339, r.FormValue("start_date"))
		finishDate, _ := time.Parse(time.RFC3339, r.FormValue("finish_date"))
		sum, _ := strconv.Atoi(r.FormValue("sum"))
		closed, _ := strconv.ParseBool(r.FormValue("closed"))
		itemsForm := r.Form["form_items"]
		allItems, _ := s.store.Item().FindAll()
		itemIds := make([]int, 0, 1)
		for _, value := range itemsForm {
			v, _ := strconv.Atoi(value)
			itemIds = append(itemIds, v)
		}

		items := s.store.Item().FindAllByID(itemIds)
		p := &model.Plan{
			ID:         id,
			Sum:        sum,
			StartDate:  startDate,
			FinishDate: finishDate,
			Closed:     closed,
			Items:      items,
			AllItems:   allItems,
		}
		if id == 0 {
			if err := s.store.Plan().Create(p); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		} else {
			p.ID = id
			if err := s.store.Plan().Update(p); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		}

		s.respondAndRedirect(w, r, http.StatusSeeOther, &p, "/plans")
	}
}

func (s *server) handleItems() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		items, err := s.store.Item().FindAll()
		if err != nil {
			log.Fatal(err)
		}
		err = s.templates.ExecuteTemplate(w, "items.gohtml", items)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) handleItemsUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		title := r.URL.Query().Get("title")
		i := &model.Item{
			ID:    id,
			Title: title,
		}

		err := s.templates.ExecuteTemplate(w, "item-form.gohtml", i)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) handleItemsUpdateProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		id, _ := strconv.Atoi(r.FormValue("id"))
		title := r.FormValue("title")

		i := &model.Item{
			Title: title,
		}
		if id == 0 {
			if err := s.store.Item().Create(i); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		} else {
			i.ID = id
			if err := s.store.Item().Update(i); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		}

		s.respondAndRedirect(w, r, http.StatusSeeOther, i, "/items")
	}
}

func (s *server) handleItemsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		s.store.Item().DeleteById(id)
		s.respondAndRedirect(w, r, http.StatusSeeOther, nil, "/items")
	}
}

func (s *server) handleInvoices() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		invoices, err := s.store.Invoice().FindAll()
		if err != nil {
			log.Fatal(err)
		}
		err = s.templates.ExecuteTemplate(w, "invoices.gohtml", invoices)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

func (s *server) respondAndRedirect(w http.ResponseWriter, r *http.Request, code int, data interface{}, redirectUrl string) {
	w.Header().Set("Location", redirectUrl)
	s.respond(w, r, code, data)
}

func (s *server) handleInvoicesDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		s.store.Invoice().DeleteById(id)
		s.respondAndRedirect(w, r, http.StatusSeeOther, nil, "/invoices")
	}
}

func (s *server) handleInvoicesUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		title := r.URL.Query().Get("title")
		description := r.URL.Query().Get("description")
		aim, _ := strconv.Atoi(r.URL.Query().Get("aim"))
		i := &model.Invoice{
			ID:          id,
			Title:       title,
			Description: description,
			Aim:         aim,
		}

		err := s.templates.ExecuteTemplate(w, "invoice-form.gohtml", i)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) handleInvoicesUpdateProcess() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		id, _ := strconv.Atoi(r.FormValue("id"))
		aim, _ := strconv.Atoi(r.FormValue("aim"))
		title := r.FormValue("title")
		description := r.FormValue("description")

		i := &model.Invoice{
			Title:       title,
			Description: description,
			Aim:         aim,
		}
		if id == 0 {
			if err := s.store.Invoice().Create(i); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		} else {
			i.ID = id
			if err := s.store.Invoice().Update(i); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		}

		s.respondAndRedirect(w, r, http.StatusSeeOther, i, "/invoices")
	}
}
