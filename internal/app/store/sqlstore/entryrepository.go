package sqlstore

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"

	"time"
)

// EntryRepository ...
type EntryRepository struct {
	store *Store
}

// Create ...
func (r *EntryRepository) Create(e *model.Entry) error {
	r.store.db.Create(&e)
	return nil
}

// FindById ...
func (r *EntryRepository) FindById(id int) (*model.Entry, error) {
	e := &model.Entry{}
	r.store.db.First(&e, id)

	return e, nil
}

// Update ...
func (r *EntryRepository) Update(e *model.Entry) error {
	r.store.db.Model(&e).Update("entry_date", e.EntryDate, "item_id", e.ItemID, "invoice_id", e.InvoiceID, "sum", e.Sum).Where("id", e.ID)
	return nil
}

// FindAllWithinPeriod ...
func (r *EntryRepository) FindAllWithinPeriod(start time.Time, finish time.Time) []*model.Entry {
	var entries []*model.Entry
	r.store.db.Where("entry_date BETWEEN ? AND ?", start, finish).Preload("Item").Preload("Invoice").Find(&entries)

	return entries
}
