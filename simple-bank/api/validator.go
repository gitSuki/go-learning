package api

import (
	"github.com/gitsuki/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldlLevel validator.FieldLevel) bool {
	if currency, ok := fieldlLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
