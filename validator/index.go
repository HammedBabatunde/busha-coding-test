package validator

import go_validator "github.com/go-playground/validator/v10"

var Validator = GoValidator{validator: go_validator.New()}
