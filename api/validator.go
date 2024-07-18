package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/baksman/backend_masterclass/util"
)

var validateCurrency validator.Func = func(fl validator.FieldLevel) bool {
	fmt.Printf("its called")
	if currency, ok := fl.Field().Interface().(string); ok {
		// check if currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
