package model

import validation "github.com/go-ozzo/ozzo-validation"

// необязательная хрень которая включает проверку на обязательность в зависимости от флажка
func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
