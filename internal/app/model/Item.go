package model

// Item ...
type Item struct {
	ID      int    `gorm:"primary_key;auto_increment", json:"id,string"` // ,string,omitempty
	Title   string `json:"title"`
	Deleted bool   `json:"deleted,string"`
}
