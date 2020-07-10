package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// Plan ...
type Plan struct {
	ID          int        `gorm:"primary_key;auto_increment" json:"id,string"`
	StartDate   time.Time  `gorm:"type:time" json:"-"`
	FinishDate  time.Time  `gorm:"type:time" json:"-"`
	StartDateS  string     `gorm:"-" json:"start_date"`
	FinishDateS string     `gorm:"-" json:"finish_date"`
	PlanItems   []PlanItem `json:"plan_items,string"` //gorm:"many2many:plan_items"
	Closed      bool       `json:"closed"`
	AllItems    []*Item    `gorm:"-" json:"-"`
}

// GetFormattedStartDate ...
func (p *Plan) GetFormattedStartDate() string {
	return p.StartDate.Format("2006-01-02")
}

// GetFormattedFinishDate ...
func (p *Plan) GetFormattedFinishDate() string {
	return p.FinishDate.Format("2006-01-02")
}

// GetPlanItemsSize ...
func (p *Plan) GetPlanItemsSize() int {
	return len(p.PlanItems)
}

func (p *Plan) UnmarshalJSON(b []byte) error {
	type Plan2 Plan

	var p2 Plan2
	err := json.Unmarshal(b, &p2)
	fmt.Println(p2)
	if err != nil {
		return err
	}
	*p = Plan(p2)
	p.FinishDate, _ = time.Parse("2006-01-02", p2.FinishDateS)
	p.StartDate, _ = time.Parse("2006-01-02", p2.StartDateS)
	return nil

}

func (p *Plan) MarshalJSON() ([]byte, error) {
	result, err := json.Marshal(p)
	return result, err
}
