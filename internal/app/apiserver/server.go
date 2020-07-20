package apiserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
	"github.com/sirupsen/logrus"
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
	s.router.HandleFunc("/", s.handleDashboard()).Methods("GET")
	s.router.HandleFunc("/entries/add", s.handleAddEntry()).Methods("POST") // todo сделай форму и аяксом ее пыщпыщ

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
	s.router.HandleFunc("/plans/update/item/add", s.handleAddItem()).Methods("POST")
	//s.router.HandleFunc("/plans/update/item/delete", s.handleDeleteItem()).Methods("POST") //todo
	//s.router.HandleFunc("/plans/update", s.handlePlansUpdate()).Methods("GET")

}

func (s *server) handleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		currentLocation := now.Location()
		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)

		var data map[string]interface{}
		data = make(map[string]interface{})

		entries := s.store.Entry().FindAllWithinPeriod(firstOfMonth, time.Now())
		data["entries"] = entries

		currentPlan, _ := s.store.Plan().FindCurrentPlan()
		data["current_plan"] = currentPlan

		allItems, _ := s.store.Item().FindAll()
		data["all_items"] = allItems

		allInvoices, _ := s.store.Invoice().FindAll()
		data["all_invoices"] = allInvoices

		err := s.templates.ExecuteTemplate(w, "dashboard.gohtml", data)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) handleAddEntry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		planItem := &model.PlanItem{}
		if err := json.Unmarshal(bs, &planItem); err != nil {
			panic(err)
		}
		item, _ := s.store.Item().FindById(planItem.ItemID)
		planItem.Item = item
		planItem.ItemTitle = item.Title
		if planItem.PlanID != 0 {
			plan, _ := s.store.Plan().FindById(planItem.PlanID)
			planItem.Plan = plan
		}
		data, err := json.Marshal(planItem)
		if err != nil {
			panic(err)
		}
		_, _ = fmt.Fprint(w, string(data))
	}
}

func (s *server) handleDeleteItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//todo если тебе очень захочется и нечего будет делать
	}
}

func (s *server) handleAddItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		planItem := &model.PlanItem{}
		if err := json.Unmarshal(bs, &planItem); err != nil {
			panic(err)
		}
		item, _ := s.store.Item().FindById(planItem.ItemID)
		planItem.Item = item
		planItem.ItemTitle = item.Title
		if planItem.PlanID != 0 {
			plan, _ := s.store.Plan().FindById(planItem.PlanID)
			planItem.Plan = plan
		}
		data, err := json.Marshal(planItem)
		if err != nil {
			panic(err)
		}
		_, _ = fmt.Fprint(w, string(data))
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
		startDateS := r.URL.Query().Get("start_date")
		startDate, _ := time.Parse("2006-01-02", startDateS)
		finishDateS := r.URL.Query().Get("finish_date")
		finishDate, _ := time.Parse("2006-01-02", finishDateS)
		closed, _ := strconv.ParseBool(r.URL.Query().Get("closed"))
		allItems, _ := s.store.Item().FindAll()

		var p *model.Plan
		if id != 0 {
			p, _ = s.store.Plan().FindById(id)
			p.AllItems = allItems
		} else {
			var planItems []*model.PlanItem
			p = &model.Plan{
				ID:          id,
				StartDate:   startDate,
				StartDateS:  startDateS,
				FinishDate:  finishDate,
				FinishDateS: finishDateS,
				Closed:      closed,
				PlanItems:   planItems,
				AllItems:    allItems,
			}
		}

		err := s.templates.ExecuteTemplate(w, "plan-form.gohtml", p)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
	}
}

func (s *server) handlePlansUpdateProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bs))
		plan := &model.Plan{}
		if err := json.Unmarshal(bs, plan); err != nil {
			fmt.Println(err)
		}
		if plan.ID != 0 {
			s.store.Plan().Update(plan)
		} else {
			_ = s.store.Plan().Create(plan)
		}
		s.respondAndRedirect(w, r, http.StatusSeeOther, nil, "/plans")
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
		_ = s.store.Item().DeleteById(id)
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
		_ = s.store.Invoice().DeleteById(id)
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
