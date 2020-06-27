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
	Create(item *model.Item) error
	FindById(int) (*model.Item, error)
	FindAll() ([]*model.Item, error)
	Update(item *model.Item) error
	DeleteById(int) error
}
