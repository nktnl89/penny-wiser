package sqlstore

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"
)

// PlanRepository ...
type PlanRepository struct {
	store *Store
}

// Create ...
func (r *PlanRepository) Create(p *model.Plan) error {
	r.store.db.Create(&p)
	return nil
}

// Update ...
func (r *PlanRepository) Update(p *model.Plan) error {
	r.store.db.Update(&p).Where("id", p.ID)
	return nil
}

// FindById ...
func (r *PlanRepository) FindById(id int) (*model.Plan, error) {
	p := &model.Plan{}
	r.store.db.First(&p, id)
	items := &[]model.Item{}
	r.store.db.Preload("items").First(&p)
	r.store.db.Model(&p).Related(&items, "Items")

	//r.store.db.Preload("items").First(&p, "id = ?", 1)
	return p, nil
}

// FindAll ...
func (r *PlanRepository) FindAll() ([]*model.Plan, error) {
	var plans []*model.Plan
	r.store.db.Find(&plans)
	return plans, nil
}

// DeleteById ...
func (r *PlanRepository) DeleteById(id int) error {
	r.store.db.Exec("with deleted_items as ("+
		"\tselect id, deleted from items where id = $1)"+
		"\tupdate items set deleted = not deleted_items.deleted"+
		"\tfrom deleted_items"+
		"\twhere items.id = deleted_items.id", id)
	return nil
}
