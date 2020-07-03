package sqlstore

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"
)

// UserRepository ...
type InvoiceRepository struct {
	store *Store
}

// Create ...
func (r *InvoiceRepository) Create(i *model.Invoice) error {
	if err := i.Validate(); err != nil {
		return err
	}
	r.store.db.Create(&i).Scan(&i)
	return nil
}

// Update ...
func (r *InvoiceRepository) Update(i *model.Invoice) error {
	if err := i.Validate(); err != nil {
		return err
	}
	r.store.db.Update(&i, "title", i.Title, "description", i.Description, "aim", i.Aim, "deleted", i.Deleted).Where(i.ID)
	return nil

}

// FindById ...
func (r *InvoiceRepository) FindById(id int) (*model.Invoice, error) {
	i := &model.Invoice{}
	r.store.db.First(&i, id)

	return i, nil
}

// FindAll ...
func (r *InvoiceRepository) FindAll() ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	r.store.db.Find(&invoices)

	return invoices, nil
}

// DeleteById ...
func (r *InvoiceRepository) DeleteById(id int) error {
	r.store.db.Exec("with deleted_invoices as ("+
		"\tselect id, deleted from invoices where id = $1)"+
		"\tupdate invoices set deleted = not deleted_invoices.deleted"+
		"\tfrom deleted_invoices"+
		"\twhere invoices.id = deleted_invoices.id", id)
	return nil
}
