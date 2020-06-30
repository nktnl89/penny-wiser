package model

import "time"

// Plan ...
type Plan struct {
	ID         int
	Sum        int
	StartDate  time.Time
	FinishDate time.Time
	Item       *Item
	Closed     bool
}
