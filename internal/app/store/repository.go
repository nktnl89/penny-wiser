package store

import "github.com/nktnl89/penny-wiser/internal/app/model"

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
}
