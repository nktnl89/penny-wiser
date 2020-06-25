package sqlstore

import (
	"database/sql"
	"github.com/nktnl89/penny-wiser/internal/app/store"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	db                *sql.DB
	invoiceRepository *InvoiceRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Invoice ...
func (s *Store) Invoice() store.InvoiceRepository {
	if s.invoiceRepository != nil {
		return s.invoiceRepository
	}
	s.invoiceRepository = &InvoiceRepository{
		store: s,
	}
	return s.invoiceRepository
}
