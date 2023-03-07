package comment

import "github.com/emekarr/coding-test-busha/validator"

type CommentPayload struct {
	Comment string `validate:"required,lt=501" json:"comment"`
}

func (cp *CommentPayload) Validate() *[]string {
	return validator.Validator.Validate(cp)
}

