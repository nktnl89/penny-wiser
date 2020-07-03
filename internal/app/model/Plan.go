package model

import "time"

// Plan ...
type Plan struct {
	ID         int `gorm:"primary_key;auto_increment"`
	Sum        int
	StartDate  time.Time
	FinishDate time.Time
	Items      []*Item `gorm:"many2many:plans_items;"`
	Closed     bool
	AllItems   []*Item `gorm:"-"`
}

// GetFormattedStartDate ...
func (p *Plan) GetFormattedStartDate() string {
	return p.StartDate.Format("02-Jan-2006")
}

// GetFormattedFinishDate ...
func (p *Plan) GetFormattedFinishDate() string {
	return p.FinishDate.Format("02-Jan-2006")
}
