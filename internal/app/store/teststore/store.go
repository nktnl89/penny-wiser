package teststore

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
)

// Store ...
type Store struct {
	invoiceRepository *InvoiceRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Invoice ...
func (s *Store) Invoice() store.InvoiceRepository {
	if s.invoiceRepository != nil {
		return s.invoiceRepository
	}
	s.invoiceRepository = &InvoiceRepository{
		store:    s,
		invoices: make(map[int]*model.Invoice),
	}
	return s.invoiceRepository
}
