package sqlstore

import (
	"database/sql"
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
	"log"
)

// ItemRepository ...
type ItemRepository struct {
	store *Store
}

// Create ...
func (r *ItemRepository) Create(i *model.Item) error {
	return r.store.db.QueryRow("INSERT INTO items (title, deleted) values ($1, false) RETURNING id",
		i.Title,
	).Scan(&i.ID)
}

// Update ...
func (r *ItemRepository) Update(i *model.Item) error {
	_, err := r.store.db.Exec("UPDATE items SET title = $2, deleted = $3 WHERE id = $1",
		i.ID,
		i.Title,
		i.Deleted)
	return err
}

// FindById ...
func (r *ItemRepository) FindById(id int) (*model.Item, error) {
	i := &model.Item{}
	if err := r.store.db.QueryRow(
		"SELECT id, title, deleted FROM items WHERE id = $1",
		id).Scan(&i.ID,
		&i.Title,
		&i.Deleted); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return i, nil
}

// FindAll ...
func (r *ItemRepository) FindAll() ([]*model.Item, error) {
	rows, err := r.store.db.Query("SELECT * FROM items order by title desc, deleted desc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var invoices []*model.Item
	for rows.Next() {
		var (
			id      int
			title   string
			deleted bool
		)
		if err := rows.Scan(&id, &title, &deleted); err != nil {
			log.Fatal(err)
		}
		invoices = append(invoices, &model.Item{
			ID:      id,
			Title:   title,
			Deleted: deleted,
		})
	}
	return invoices, nil
}

// DeleteById ...
func (r *ItemRepository) DeleteById(id int) error {
	_, err := r.store.db.Exec("with deleted_items as ("+
		"\tselect id, deleted from items where id = $1)"+
		"\tupdate items set deleted = not deleted_items.deleted"+
		"\tfrom deleted_items"+
		"\twhere items.id = deleted_items.id", id)
	return err
}
