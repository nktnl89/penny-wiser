package teststore

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
)

type InvoiceRepository struct {
	store    *Store
	invoices map[int]*model.Invoice
}

func (r *InvoiceRepository) Create(i *model.Invoice) error {
	if err := i.Validate(); err != nil {
		return err
	}
	r.invoices[i.ID] = i
	i.ID = len(r.invoices)

	return nil
}

func (r *InvoiceRepository) FindById(id int) (*model.Invoice, error) {
	i, ok := r.invoices[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return i, nil
}

// FindAll ...
func (u *InvoiceRepository) FindAll() ([]*model.Invoice, error) {

	// m := []*model.Invoice{
	// 	{0, "cash", "", 0},
	// 	{1, "alfa", "", 0},
	// }
	// v := make([]*model.Invoice, 0, len(m))

	// for _, value := range m {
	// 	v = append(v, value)
	// }

	return nil, nil
}
