package entities

// Invoice ...
type Invoice struct {
	ID    int
	Title string
	Plan  int
}

// HasPlan ...
func (i Invoice) HasPlan() bool {
	return i.Plan > 0
}

// GetCurrentSum ...
func (i Invoice) GetCurrentSum() int {
	return 100 // надо прикрутить расчет текущего значения по таблице с items
}
