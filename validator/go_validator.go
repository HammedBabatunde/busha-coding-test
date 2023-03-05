package validator

import go_validator "github.com/go-playground/validator/v10"

type GoValidator struct {
	validator *go_validator.Validate
}

func (g *GoValidator) Validate(payload interface{}) error {
	return g.validator.Struct(payload)
}
