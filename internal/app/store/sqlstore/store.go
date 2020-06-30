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
	itemRepository    *ItemRepository
	planRepository    *PlanRepository
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

// Item ...
func (s *Store) Item() store.ItemRepository {
	if s.itemRepository != nil {
		return s.itemRepository
	}
	s.itemRepository = &ItemRepository{
		store: s,
	}
	return s.itemRepository
}

// Item ...
func (s *Store) Plan() store.PlanRepository {
	if s.planRepository != nil {
		return s.planRepository
	}
	s.planRepository = &PlanRepository{
		store: s,
	}
	return s.planRepository
}
