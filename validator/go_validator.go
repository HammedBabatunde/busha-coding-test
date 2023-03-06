package validator

import (
	"strings"

	go_validator "github.com/go-playground/validator/v10"
)

type GoValidator struct {
	validator *go_validator.Validate
}

func (g *GoValidator) ParseErrorMessage(msg string) *[]string {
	errs := strings.Split(msg, "Key: ")
	errsParsed := []string{}
	for _, e := range errs {
		errsParsed = append(errsParsed, strings.TrimSpace(strings.ReplaceAll(e, " tag", "")))
	}

	errsParsed = errsParsed[1:]
	return &errsParsed
}

func (g *GoValidator) Validate(payload interface{}) *[]string {
	err := g.validator.Struct(payload)
	if err != nil {
		return g.ParseErrorMessage(err.Error())
	}
	return nil
}
