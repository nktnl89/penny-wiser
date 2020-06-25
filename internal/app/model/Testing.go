package model

import "testing"

// TestInvoice ...
func TestInvoice(t *testing.T) *Invoice {
	return &Invoice{
		Title:       "test",
		Description: "test description",
		Aim:         100000,
	}
}
