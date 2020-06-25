package sqlstore

import (
	"database/sql"
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
	"log"
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
	return r.store.db.QueryRow("INSERT INTO invoices (title, description, aim) values ($1, $2, $3) RETURNING id",
		i.Title,
		i.Description,
		i.Aim,
	).Scan(&i.ID)
}

// Update ...
func (r *InvoiceRepository) Update(i *model.Invoice) error {
	if err := i.Validate(); err != nil {
		return err
	}
	_, err := r.store.db.Exec("UPDATE invoices SET title = $2, description = $3, aim = $4 WHERE id = $1",
		i.ID,
		i.Title,
		i.Description,
		i.Aim)

	return err
}

// FindById ...
func (r *InvoiceRepository) FindById(id int) (*model.Invoice, error) {
	i := &model.Invoice{}
	//defer r.store.db.Close()
	if err := r.store.db.QueryRow(
		"SELECT id, title, description, aim FROM invoices WHERE id = $1",
		id).Scan(&i.ID,
		&i.Title,
		&i.Description,
		&i.Aim); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return i, nil
}

// FindAll ...
func (r *InvoiceRepository) FindAll() ([]*model.Invoice, error) {
	rows, err := r.store.db.Query("SELECT * FROM invoices WHERE deleted = false")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var invoices []*model.Invoice
	for rows.Next() {
		var (
			id          int
			title       string
			description string
			aim         int
		)
		if err := rows.Scan(&id, &title, &description, &aim); err != nil {
			log.Fatal(err)
		}
		invoices = append(invoices, &model.Invoice{
			ID:          id,
			Title:       title,
			Description: description,
			Aim:         aim,
		})
	}
	return invoices, nil
}
