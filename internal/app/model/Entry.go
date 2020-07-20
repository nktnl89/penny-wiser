package model

import (
	"time"
)

// Entry ...
type Entry struct {
	ID         int       `gorm:"primary_key;auto_increment" json:"id,string"`
	EntryDate  time.Time `gorm:"type:time" json:"-"`
	EntryDateS string    `gorm:"-" json:"entry_date"`
	ItemID     int       `json:"item_id,string"`
	Item       *Item     `gorm:"foreignkey:ItemID" json:"-"`
	InvoiceID  int       `json:"invoice_id,string"`
	Invoice    *Invoice  `gorm:"foreignkey:InvoiceID" json:"-"`
	Sum        int       `json:"sum"`
}

// GetFormattedDate ...
func (e *Entry) GetFormattedDate() string {
	return e.EntryDate.Format("2006-01-02")
}
