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
