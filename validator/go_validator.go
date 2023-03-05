package validator

import (
	"strings"

	go_validator "github.com/go-playground/validator/v10"
)

type GoValidator struct {
	validator *go_validator.Validate
}

func (g *GoValidator) ParseErrorMessage(msg string) *[]string {
	errs := strings.Split(msg, "Key")
	return &errs
}

func (g *GoValidator) Validate(payload interface{}) *[]string {
	return g.ParseErrorMessage(g.validator.Struct(payload).Error())
}
