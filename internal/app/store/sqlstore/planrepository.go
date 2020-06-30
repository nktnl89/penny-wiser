package sqlstore

import (
	"database/sql"
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
	"log"
	"time"
)

// PlanRepository ...
type PlanRepository struct {
	store *Store
}

// Create ...
func (r *PlanRepository) Create(p *model.Plan) error {
	return r.store.db.QueryRow("INSERT INTO plans (item_id, start_date, finish_date, sum, closed) "+
		"	values ((select i.id from items i where i.title = $1), $2, $3, $4, $5) RETURNING id",
		p.Item.Title, p.StartDate, p.FinishDate, p.Sum, false,
	).Scan(&p.ID)
}

// Update ...
func (r *PlanRepository) Update(p *model.Plan) error {
	_, err := r.store.db.Exec("UPDATE plans SET item_id = $2, start_date = $3, finish_date = $4, sum = $5, closed = $6 WHERE id = $1",
		p.ID,
		p.Item.ID,
		p.StartDate,
		p.FinishDate,
		p.Sum,
		p.Closed)
	return err
}

// FindById ...
func (r *PlanRepository) FindById(id int) (*model.Plan, error) {
	p := &model.Plan{}
	i := &model.Item{}
	if err := r.store.db.QueryRow(
		"SELECT p.id, p.start_date, p.finish_date, p.sum, p.closed, i.id, i.title, i.deleted FROM plans as p inner join items as i on p.item_id = i.id WHERE p.id = $1", id).Scan(&p.ID,
		&p.StartDate,
		&p.FinishDate,
		&p.Sum,
		&p.Closed,
		&i.ID,
		&i.Title,
		&i.Deleted); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	p.Item = i
	return p, nil
}

// FindAll ...
func (r *PlanRepository) FindAll() ([]*model.Plan, error) {
	rows, err := r.store.db.Query("SELECT p.id, p.start_date, p.finish_date, p.sum, p.closed, i.id, i.title, i.deleted FROM plans as p inner join items as i on p.item_id = i.id order by p.id desc, p.closed desc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var plans []*model.Plan
	for rows.Next() {
		var (
			id          int
			closed      bool
			startDate   time.Time
			finishDate  time.Time
			sum         int
			itemId      int
			itemTitle   string
			itemDeleted bool
		)
		if err := rows.Scan(&id, &startDate, &finishDate, &sum, &closed, &itemId, &itemTitle, &itemDeleted); err != nil {
			log.Fatal(err)
		}
		plans = append(plans, &model.Plan{
			ID:         id,
			StartDate:  startDate,
			FinishDate: finishDate,
			Sum:        sum,
			Closed:     closed,
			Item: &model.Item{
				ID:      itemId,
				Title:   itemTitle,
				Deleted: itemDeleted,
			},
		})
	}
	return plans, nil
}

// DeleteById ...
func (r *PlanRepository) DeleteById(id int) error {
	_, err := r.store.db.Exec("with deleted_items as ("+
		"\tselect id, deleted from items where id = $1)"+
		"\tupdate items set deleted = not deleted_items.deleted"+
		"\tfrom deleted_items"+
		"\twhere items.id = deleted_items.id", id)
	return err
}
