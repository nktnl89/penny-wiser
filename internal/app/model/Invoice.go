package model

import validation "github.com/go-ozzo/ozzo-validation"

// Invoice ...
type Invoice struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Aim         int    `json:"aim,omitempty"`
}

// HasPlan ...
func (i *Invoice) HasAim() bool {
	return i.Aim > 0
}

// GetCurrentSum ...
func (i *Invoice) GetCurrentSum() int {
	return 100 // надо прикрутить расчет текущего значения по таблице с items
}

// Validate ...
func (i *Invoice) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Title, validation.Required),
		validation.Field(&i.Description, validation.Length(0, 100)),
		validation.Field(&i.Aim, validation.Min(0)),
	)
}
