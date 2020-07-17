package model

// Item ...
type Item struct {
	ID      int    `gorm:"primary_key auto_increment" json:"id,string"`
	Title   string `gorm:"type:string" json:"title"`
	Deleted bool   `gorm:"type:bool" json:"deleted,string"`
}
