package store

import (
	"time"

	"github.com/nktnl89/penny-wiser/internal/app/model"
)

// InvoiceRepository ...
type InvoiceRepository interface {
	Create(*model.Invoice) error
	FindById(int) (*model.Invoice, error)
	FindAll() ([]*model.Invoice, error)
	Update(*model.Invoice) error
	DeleteById(int) error
}

// ItemRepository ...
type ItemRepository interface {
	Create(*model.Item) error
	FindById(int) (*model.Item, error)
	FindAll() ([]*model.Item, error)
	Update(item *model.Item) error
	DeleteById(int) error
	FindAllByID([]int) []*model.Item
}

// PlanRepository ...
type PlanRepository interface {
	Create(*model.Plan) error
	FindById(int) (*model.Plan, error)
	FindAll() ([]*model.Plan, error)
	Update(*model.Plan) error
	DeleteById(int) error
	FindCurrentPlan() (*model.Plan, error)
}

// PlanItemRepository ...
type PlanItemRepository interface {
	Create(item *model.PlanItem) error
	FindById(int) (*model.PlanItem, error)
	Update(*model.PlanItem) error
	FindAllByPlanID(int) []*model.PlanItem
}

// EntryRepository ...
type EntryRepository interface {
	Create(e *model.Entry) error
	FindById(id int) (*model.Entry, error)
	Update(e *model.Entry) error
	FindAllWithinPeriod(start time.Time, finish time.Time) []*model.Entry
}
