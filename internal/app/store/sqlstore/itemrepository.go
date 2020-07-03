package sqlstore

import (
	"fmt"
	"github.com/nktnl89/penny-wiser/internal/app/model"
)

// ItemRepository ...
type ItemRepository struct {
	store *Store
}

// Create ...
func (r *ItemRepository) Create(i *model.Item) error {
	//r.store.db.AutoMigrate(&model.Item{})
	r.store.db.Create(&i)
	return nil
}

// Update ...
func (r *ItemRepository) Update(i *model.Item) error {
	r.store.db.Model(&i).Update("title", i.Title, "deleted", i.Deleted).Where("id", i.ID)
	return nil
}

// FindById ...
func (r *ItemRepository) FindById(id int) (*model.Item, error) {
	i := &model.Item{}
	r.store.db.First(&i, id)
	return i, nil
}

// FindAll ...
func (r *ItemRepository) FindAll() ([]*model.Item, error) {
	var items []*model.Item
	r.store.db.Find(&items)
	return items, nil
}

// DeleteById ...
func (r *ItemRepository) DeleteById(id int) error {
	r.store.db.Exec("with deleted_items as ("+
		"\tselect id, deleted from items where id = $1)"+
		"\tupdate items set deleted = not deleted_items.deleted"+
		"\tfrom deleted_items"+
		"\twhere items.id = deleted_items.id", id)
	return nil
}

// FindAllByID ...
func (r *ItemRepository) FindAllByID(ids []int) []*model.Item {
	var items []*model.Item
	for v, _ := range ids {
		fmt.Println(v)
	}
	if len(ids) == 0 {
		return items
	}
	r.store.db.Where("id IN ?", ids).Find(&items)
	return items
}
