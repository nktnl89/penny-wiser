package model

// PlanItem ...
type PlanItem struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"-"` //id,string
	PlanID    int    `json:"plan_id,string"`
	ItemID    int    `json:"item_id,string"`
	ItemTitle string `gorm:"-" json:"item_title"`
	Sum       int    `json:"sum,string"`
	Plan      *Plan  `gorm:"foreignkey:PlanID" json:"-"`
	Item      *Item  `gorm:"foreignkey:ItemID" json:"-"`
}
