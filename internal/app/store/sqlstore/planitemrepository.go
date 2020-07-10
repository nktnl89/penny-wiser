package sqlstore

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"
)

// PlanItemRepository ...
type PlanItemRepository struct {
	store *Store
}

// Create ...
func (r *PlanItemRepository) Create(pi *model.PlanItem) error {
	r.store.db.Create(&pi)
	return nil
}

// Update ...
func (r *PlanItemRepository) Update(pi *model.PlanItem) error {
	r.store.db.Model(&pi).Update("sum", pi.Sum, "plan_id", pi.PlanID, "item_id", pi.ItemID).Where("id", pi.ID)
	return nil
}

// FindById ...
func (r *PlanItemRepository) FindById(id int) (*model.PlanItem, error) {
	pi := &model.PlanItem{}
	r.store.db.First(&pi, id)
	return pi, nil
}

// FindAllByPlanID ...
func (r *PlanItemRepository) FindAllByPlanID(id int) []*model.PlanItem {
	var planItems []*model.PlanItem

	r.store.db.Where("plan_id = ?", id).Find(&planItems)
	return planItems
}
