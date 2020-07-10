package model

// PlanItem ...
type PlanItem struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"-"` //id,string
	PlanID    int    `gorm:"foreignkey:PlanID;association_foreignkey:Plan" json:"plan_id,string"`
	ItemID    int    `gorm:"foreignkey:ItemID;association_foreignkey:Item" json:"item_id,string"`
	ItemTitle string `gorm:"-" json:"item_title"`
	Sum       int    `json:"sum,string"`
	Plan      *Plan  `json:"-"`
	Item      *Item  `json:"-"`
}
