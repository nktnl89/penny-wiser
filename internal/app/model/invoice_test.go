package model_test

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

//todo надо б дописать сюда тест по расчету текущего потраченого со счета
//todo и тест по поводу хэзЭйм

func TestInvoice_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		i       func() *model.Invoice
		isValid bool
	}{
		{
			name: "valid",
			i: func() *model.Invoice {
				return model.TestInvoice(t)
			},
			isValid: true,
		},
		{
			name: "negative aim",
			i: func() *model.Invoice {
				inv := model.TestInvoice(t)
				inv.Aim = -10
				return inv
			},
			isValid: false,
		},
		{
			name: "valid",
			i: func() *model.Invoice {
				inv := model.TestInvoice(t)
				inv.Aim = 40000
				return inv
			},
			isValid: true,
		},
		{
			name: "too long description",
			i: func() *model.Invoice {
				inv := model.TestInvoice(t)
				inv.Description = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111" +
					"111111111111111111111111111111111111111111111"
				return inv
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.i().Validate())
			} else {
				assert.Error(t, tc.i().Validate())
			}
		})
	}
}
