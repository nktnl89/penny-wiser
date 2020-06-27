package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
	"github.com/sirupsen/logrus"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	s.router.HandleFunc("/periods", s.handlePeriods()).Methods("GET")
	s.router.HandleFunc("/periods/update", s.handlePeriodsUpdate()).Methods("GET")
	s.router.HandleFunc("/periods/update/process", s.handlePeriodsUpdateProcess()).Methods("POST")
}

func (s *server) handleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

	}
}

func (s *server) handlePeriods() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

	}
}

func (s *server) handlePeriodsUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

	}
}

func (s *server) handlePeriodsUpdateProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

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

//func (s *APIServer) configureLogger() error {
//	level, err := logrus.ParseLevel(s.config.LogLevel)
//
//	if err != nil {
//		return err
//	}
//	s.logger.SetLevel(level)
//	return nil
//}
//
//func (s *APIServer) configureRouter() {
//	s.router.HandleFunc("/hello", s.handleHello())
//	s.router.HandleFunc("/", s.handleDashboard())
//	s.router.HandleFunc("/invoices", s.handleInvoices())
//	s.router.HandleFunc("/items", s.handleItems())
//	//s.router.HandleFunc("/items/add", s.handleItemsAdding())
//	//s.router.HandleFunc("/plans", s.handlePlans())
//	//s.router.HandleFunc("/periods", s.handlePeriods())
//}
//
//func (s *APIServer) handleHello() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		_, _ = io.WriteString(writer, "Hello")
//	}
//}
//
//func (s *APIServer) handleDashboard() http.HandlerFunc {
//	items := []model.Item{
//		{0, "Кредит"},
//		{1, "Кварплата"},
//	}
//	return func(w http.ResponseWriter, request *http.Request) {
//		err := tpl.ExecuteTemplate(w, "items.gohtml", items)
//		if err != nil {
//			log.Fatalln("template didn't execute: ", err)
//		}
//	}
//}
//
//func (s *APIServer) handleItems() http.HandlerFunc {
//	items := []model.Item{
//		{0, "Кредит"},
//		{1, "Кварплата"},
//	}
//	return func(w http.ResponseWriter, request *http.Request) {
//		err := tpl.ExecuteTemplate(w, "items.gohtml", items)
//		if err != nil {
//			log.Fatalln("template didn't execute: ", err)
//		}
//	}
//}
//
//func (s *APIServer) itemsAdding() http.HandlerFunc {
//	//items := []Item{Item{0, "Кредит"}, Item{1, "Кварплата"}}
//	//
//	//err := tpl.ExecuteTemplate(res, "items.gohtml", items)
//	//if err != nil {
//	//	log.Fatalln("template didn't execute: ", err)
//	//}
//}
//func (s *APIServer) plans() http.HandlerFunc {
//	//items := []Item{Item{0, "Кредит"}, Item{1, "Кварплата"}}
//	//
//	//err := tpl.ExecuteTemplate(res, "items.gohtml", items)
//	//if err != nil {
//	//	log.Fatalln("template didn't execute: ", err)
//	//}
//}
//func (s *APIServer) periods() http.HandlerFunc {
//	//items := []Item{Item{0, "Кредит"}, Item{1, "Кварплата"}}
//	//
//	//err := tpl.ExecuteTemplate(res, "items.gohtml", items)
//	//if err != nil {
//	//	log.Fatalln("template didn't execute: ", err)
//	//}
//}
//
