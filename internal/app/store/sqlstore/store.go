package sqlstore

import (
	"github.com/jinzhu/gorm"
	"github.com/nktnl89/penny-wiser/internal/app/store"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Store ...
type Store struct {
	db                 *gorm.DB
	invoiceRepository  *InvoiceRepository
	itemRepository     *ItemRepository
	planRepository     *PlanRepository
	planItemRepository *PlanItemRepository
	entryRepository    *EntryRepository
}

// New ...
func New(db *gorm.DB) *Store {
	db.LogMode(false)
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

// Plan ...
func (s *Store) Plan() store.PlanRepository {
	if s.planRepository != nil {
		return s.planRepository
	}
	s.planRepository = &PlanRepository{
		store: s,
	}
	return s.planRepository
}

// PlanItem ...
func (s *Store) PlanItem() store.PlanItemRepository {
	if s.planItemRepository != nil {
		return s.planItemRepository
	}
	s.planItemRepository = &PlanItemRepository{
		store: s,
	}
	return s.planItemRepository
}

// Entry ...
func (s *Store) Entry() store.EntryRepository {
	if s.planItemRepository != nil {
		return s.entryRepository
	}
	s.entryRepository = &EntryRepository{
		store: s,
	}
	return s.entryRepository
}
